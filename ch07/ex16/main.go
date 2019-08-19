package main

// note
// POSTのパラメータにeqを入れるとeqという変数名が使えなくなる
// 一時的にidを発行して、そこにeqを格納する方針にする
// 何回も計算し直しできるし

import (
	"fmt"
	"go_training/ch07/ex15/eval"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

func envInput(vars map[eval.Var]bool) (env eval.Env) {
	env = eval.Env{}

	for v := range vars {
		fmt.Printf("%s: ", v)
		var valueStr string
		fmt.Scanln(&valueStr)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		env[v] = value
		break
	}
	return env
}

func getVars(input string) (_ []eval.Var, e error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
			e = fmt.Errorf("%s", p)
		}
	}()

	expr, err := eval.Parse(input)
	if err != nil {
		fmt.Println(err)
	}
	vars := map[eval.Var]bool{}
	expr.Check(vars)

	res := []eval.Var{}
	for v := range vars {
		res = append(res, v)
	}
	return res, nil
}

type EqInputForm struct {
	ID string
}

var eqInputForm = template.Must(template.New("eqInputForm").Parse(`
<html>
<head>
  <meta charset="UTF-8" />
</head>
<body>
<div>
  <form method="POST" action="/input?id={{.ID}}">     
      <label>eq: </label><input name="eq" type="text" value="" />
      <input type="submit" value="submit" />
  </form>
</div>
</body>
</html>

`))

type VarInputForm struct {
	EQ   string
	ID   string
	Vars []eval.Var
}

var varInputForm = template.Must(template.New("varInputForm").Parse(`
<html>
<head>
  <meta charset="UTF-8" />
</head>
<body>
<div>
  <form method="POST" action="/calc?id={{.ID}}">
	  <label>eq: {{.EQ}} </label> <br>
      {{range .Vars}}
	  <label>{{.}}: </label><input name="{{.}}" type="text" value="" /> <br>
	  {{end}}
      <input type="submit" value="calc" />
	<a href="/input">home</a>
  </form>
</div>
</body>
</html>
`))

type AnsForm struct {
	EQ        string
	ID        string
	Ans       string
	Err       string
	VarValues map[eval.Var]string
}

var ansForm = template.Must(template.New("ansForm").Parse(`
<html>
<head>
  <meta charset="UTF-8" />
</head>
<body>
<div>
  <form method="POST" action="/calc?id={{.ID}}">
	  eq: {{.EQ}} <br>
      {{range $var, $value := .VarValues}}
	  <label>{{$var}}: </label><input name="{{$var}}" type="text" value="{{$value}}" /> <br>
	  {{end}}
	  {{ if ne .Err "" }}
	  ans: {{.Err}} <br>
	  {{else}}
	  ans: {{.Ans}} <br>
	  {{end}}
	  <input type="submit" value="calc" />
  </form>
  <a href="/input">home</a>
</div>
</body>
</html>
`))

func GetRandomStrng() string {
	return strconv.FormatUint(rand.Uint64(), 10)
}

var mu sync.Mutex

var eqMap = map[string]string{} // id: eq

func main() {
	// evalInput()
	http.HandleFunc("/input", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET": // 初期画面
			if err := eqInputForm.Execute(w, EqInputForm{GetRandomStrng()}); err != nil {
				fmt.Fprintf(w, "parse error")
				log.Printf("PraseForm() err: %v", err)
				return
			}
		case "POST": // IDをsetするためのMethod
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "parse error")
				log.Printf("PraseForm() err: %v", err)
				break
			}

			ids, ok := req.Form["id"]
			if !ok {
				fmt.Fprintf(w, "id not found")
				log.Printf("id not found")
				return
			}
			if len(ids) < 1 {
				fmt.Fprintf(w, "something wrong...")
				log.Print("len(eq) < 1 error")
				return
			}
			id := ids[0]

			eqs, ok := req.PostForm["eq"]
			if !ok {
				fmt.Fprintf(w, "something wrong...")
				log.Print("req.PostForm['eq'] not found")
				return
			}
			if len(eqs) < 1 {
				fmt.Fprintf(w, "something wrong...")
				log.Print("len(eq) < 1 error")
				return
			}
			eq := eqs[0]

			mu.Lock()
			eqMap[id] = eq
			mu.Unlock()

			vars, err := getVars(eq)
			if err != nil {
				http.Redirect(w, req, "/input", 301)
				break
			}

			vif := VarInputForm{eq, id, vars}
			if err := varInputForm.Execute(w, vif); err != nil {
				fmt.Fprintf(w, "Oops something wrong...")
				log.Print(err)
				return
			}
		}
	})

	http.HandleFunc("/calc", func(w http.ResponseWriter, req *http.Request) { // idの式をいろいろするためのmethod
		defer func() {
			if p := recover(); p != nil {
				fmt.Println(p)
			}
		}()

		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "parse error")
			log.Printf("PraseForm() err: %v", err)
			return
		}

		ids, ok := req.Form["id"]
		if !ok {
			fmt.Fprintf(w, "id not found")
			log.Printf("id not found")
			return
		}
		if len(ids) < 1 {
			fmt.Fprintf(w, "something wrong...")
			log.Print("len(eq) < 1 error")
			return
		}
		id := ids[0]

		eq := eqMap[id]

		// 式の解析
		expr, err := eval.Parse(eq)
		if err != nil {
			fmt.Println("parse error")
			log.Printf("%s", err)
			return
		}

		vars := map[eval.Var]bool{}
		expr.Check(vars)

		// vars check
		if len(req.PostForm) != len(vars) {
			fmt.Fprintf(w, "form error")
			log.Print("len(req.PostForm) != len(vars)")
			return
		}

		env := eval.Env{}

		// 変数の解析
		errMessage := ""
		varValues := map[eval.Var]string{}
		for vString, valueStrs := range req.PostForm {
			v := eval.Var(vString)

			if _, ok := vars[v]; !ok {
				errMessage = fmt.Sprintf("analyze var error")
				log.Printf("%v != %v", vars, req.PostForm)
				continue
			}
			if len(valueStrs) < 1 {
				errMessage = fmt.Sprintf("something wrong...")
				log.Print("len(valueStr) < 1 error")
				continue
			}
			valueStr := valueStrs[0]

			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				errMessage = fmt.Sprintf("%s cannot parse to float!", valueStr)
				log.Print(w, "%s cannot parse to float", valueStr)
				value = 0.0
			}
			env[v] = value
			varValues[v] = strconv.FormatFloat(value, 'f', 6, 64)
		}

		ans := strconv.FormatFloat(expr.Eval(env), 'G', 6, 64)
		if err := ansForm.Execute(w, AnsForm{eq, id, ans, errMessage, varValues}); err != nil {
			fmt.Fprintf(w, "Oops something wrong...")
			log.Print(err)
			return
		}
	})

	fmt.Println("http://localhost:8000/input")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
