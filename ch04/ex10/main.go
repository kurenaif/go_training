package main

import (
	"fmt"
	"go_training/ch04/ex10/github"
	"log"
	"os"
	"time"
)

type During int

const (
	OneMonth During = iota
	OneYear
	Past
)

func (c During) String() string {
	switch c {
	case OneMonth:
		return "OneMonth"
	case OneYear:
		return "OneYear"
	case Past:
		return "Past"
	default:
		return "Unknown"
	}
}

func main() {
	cnt := make(map[During]int)
	now := time.Now()
	aMonthAgo := now.AddDate(0, -1, 0)
	aYearAGo := now.AddDate(-1, 0, 0)

	for page := 0; page < 10; page++ {
		result, err := github.SearchIssues(os.Args[1], os.Args[2:], page)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
			if aYearAGo.After(item.CreatedAt) || aYearAGo.Equal(item.CreatedAt) { // 1年以上前
				cnt[Past]++
			} else if aMonthAgo.After(item.CreatedAt) || aMonthAgo.Equal(item.CreatedAt) { //1月以上1年未満
				cnt[OneYear]++
			} else { // 1月未満
				cnt[OneMonth]++
			}
		}
	}

	fmt.Println("分類結果")
	fmt.Println("1ヶ月未満", cnt[OneMonth])
	fmt.Println("1年未満", cnt[OneYear])
	fmt.Println("1年以上", cnt[Past])
}

/*
//!+textoutput
$ go run main.go $token  https://github.com/rust-lang/rust
47295 issues:
#1753    taiki-e Elide lifetimes in `Pin<&mut Self>` 2019-07-20 09:43:16 +0000 UTC
#62727 SimonSapi Deprecate using rustc_plugin without the rustc_driver d 2019-07-16 17:46:31 +0000 UTC
#782   jonas-sch Adjust #[doc(include)] paths for rustdoc change 2019-07-17 21:03:07 +0000 UTC
#1433        jdm Link to Windows debug CRTs when crt-debug target_featur 2019-07-09 15:50:32 +0000 UTC
#62816  estebank Point at type ascription before macro invocation on exp 2019-07-20 01:14:18 +0000 UTC
#62805   Xanewok Update RLS 2019-07-19 16:05:57 +0000 UTC
#837    RalfJung test arrray try_from (interesting const generic usage) 2019-07-13 08:51:25 +0000 UTC
#385   Mark-Simu Migration fixes 2019-07-22 14:10:25 +0000 UTC
#62853  fakenine normalize use of backticks in compiler messages for lib 2019-07-21 14:56:37 +0000 UTC
#62832  fakenine normalize use of backticks in compiler messages for lib 2019-07-20 17:24:44 +0000 UTC
#62801    bjorn3 Remove support for -Zlower-128bit-ops 2019-07-19 14:05:24 +0000 UTC
#62810  fakenine normalize use of backticks in compiler messages for lib 2019-07-19 18:06:08 +0000 UTC
#62812  fakenine normalize use of backticks in compiler messages for lib 2019-07-19 19:25:50 +0000 UTC
#62788  fakenine normalize use of backticks in compiler messages for lib 2019-07-18 21:33:57 +0000 UTC
#62746  RalfJung  do not use assume_init in std::io 2019-07-17 07:54:37 +0000 UTC
#3686  topecongi New syntax in rustfmt 1.4.0 2019-07-14 14:29:07 +0000 UTC
#62656  RalfJung explain how to search in slice without owned data 2019-07-13 13:16:47 +0000 UTC
#631       ehuss Document type_alias_enum_variants 2019-07-05 16:21:48 +0000 UTC
#1757     SrTobi Remove mentioning of default spawners from docs 2019-07-21 08:10:39 +0000 UTC
#62823  RalfJung update Miri 2019-07-20 11:44:32 +0000 UTC
#2032   emmericp ci: validate that all used references are defined 2019-07-14 20:30:06 +0000 UTC
#635   nikomatsa async-await initial reference material 2019-07-10 21:27:03 +0000 UTC
#62587 SimonSapi Add serde to the `cargotest` test suite 2019-07-11 14:11:03 +0000 UTC
#62        ehuss Swap unsafe/async 2019-06-29 15:52:32 +0000 UTC
#23      taiki-e Support trailing commas in pin_mut! macro 2019-06-26 17:32:44 +0000 UTC
#2002  willstott [Rust] Add u128 and i128 primitive types. 2019-06-27 20:25:16 +0000 UTC
#62713 SimonSapi Stabilize <*mut _>::cast and <*const _>::cast 2019-07-16 06:54:56 +0000 UTC
#43        ehuss Add unstable const generics 2019-06-28 02:49:48 +0000 UTC
#45     GrayJack Parse new macro declarative expression 2019-06-26 14:34:32 +0000 UTC
#4        pyrrho Create a hub.docker.com webhook to rebuild when a new v 2019-07-04 19:31:31 +0000 UTC
47295 issues:
#1753    taiki-e Elide lifetimes in `Pin<&mut Self>` 2019-07-20 09:43:16 +0000 UTC
#62727 SimonSapi Deprecate using rustc_plugin without the rustc_driver d 2019-07-16 17:46:31 +0000 UTC
#782   jonas-sch Adjust #[doc(include)] paths for rustdoc change 2019-07-17 21:03:07 +0000 UTC
#1433        jdm Link to Windows debug CRTs when crt-debug target_featur 2019-07-09 15:50:32 +0000 UTC
#62816  estebank Point at type ascription before macro invocation on exp 2019-07-20 01:14:18 +0000 UTC
#62805   Xanewok Update RLS 2019-07-19 16:05:57 +0000 UTC
#837    RalfJung test arrray try_from (interesting const generic usage) 2019-07-13 08:51:25 +0000 UTC
#385   Mark-Simu Migration fixes 2019-07-22 14:10:25 +0000 UTC
#62853  fakenine normalize use of backticks in compiler messages for lib 2019-07-21 14:56:37 +0000 UTC
#62832  fakenine normalize use of backticks in compiler messages for lib 2019-07-20 17:24:44 +0000 UTC
#62801    bjorn3 Remove support for -Zlower-128bit-ops 2019-07-19 14:05:24 +0000 UTC
#62810  fakenine normalize use of backticks in compiler messages for lib 2019-07-19 18:06:08 +0000 UTC
#62812  fakenine normalize use of backticks in compiler messages for lib 2019-07-19 19:25:50 +0000 UTC
#62788  fakenine normalize use of backticks in compiler messages for lib 2019-07-18 21:33:57 +0000 UTC
#62746  RalfJung  do not use assume_init in std::io 2019-07-17 07:54:37 +0000 UTC
#3686  topecongi New syntax in rustfmt 1.4.0 2019-07-14 14:29:07 +0000 UTC
#62656  RalfJung explain how to search in slice without owned data 2019-07-13 13:16:47 +0000 UTC
#631       ehuss Document type_alias_enum_variants 2019-07-05 16:21:48 +0000 UTC
#1757     SrTobi Remove mentioning of default spawners from docs 2019-07-21 08:10:39 +0000 UTC
#62823  RalfJung update Miri 2019-07-20 11:44:32 +0000 UTC
#2032   emmericp ci: validate that all used references are defined 2019-07-14 20:30:06 +0000 UTC
#635   nikomatsa async-await initial reference material 2019-07-10 21:27:03 +0000 UTC
#62587 SimonSapi Add serde to the `cargotest` test suite 2019-07-11 14:11:03 +0000 UTC
#62        ehuss Swap unsafe/async 2019-06-29 15:52:32 +0000 UTC
#23      taiki-e Support trailing commas in pin_mut! macro 2019-06-26 17:32:44 +0000 UTC
#2002  willstott [Rust] Add u128 and i128 primitive types. 2019-06-27 20:25:16 +0000 UTC
#62713 SimonSapi Stabilize <*mut _>::cast and <*const _>::cast 2019-07-16 06:54:56 +0000 UTC
#43        ehuss Add unstable const generics 2019-06-28 02:49:48 +0000 UTC
#45     GrayJack Parse new macro declarative expression 2019-06-26 14:34:32 +0000 UTC
#4        pyrrho Create a hub.docker.com webhook to rebuild when a new v 2019-07-04 19:31:31 +0000 UTC
47295 issues:
#62038      Zoxc [WIP] Make dep node indices persistent between sessions 2019-06-21 21:43:16 +0000 UTC
#1          Erk- Fix compilation 2019-07-13 08:53:19 +0000 UTC
#62735 petrochen Turn `#[global_allocator]` into a regular attribute mac 2019-07-16 22:52:03 +0000 UTC
#62507 petrochen [WIP] Feature gate `(Rustc){En,De}codable` and `bench`  2019-07-08 21:36:13 +0000 UTC
#169    RalfJung Specification of interrupts, interrupt handlers, signal 2019-07-18 14:46:12 +0000 UTC
#2108  JohnTitor Replace ONCE_INIT with Once::new() 2019-07-07 01:22:16 +0000 UTC
#63        ehuss Arbitrary enum discriminants are nightly only 2019-07-10 19:10:55 +0000 UTC
#1510    matklad Resolve out-of-line modules, declared inside inline mod 2019-07-08 09:09:23 +0000 UTC
#137    tmccombs Add support for Key-Value data in log records. 2019-07-02 00:04:14 +0000 UTC
#3926  mchernyav RUN: Support build tool window 2019-05-30 23:18:19 +0000 UTC
#6       dtolnay No longer compiles 2019-07-07 07:02:14 +0000 UTC
#132     gakonst Remove `extern crate` directives since we are in Rust 2 2019-07-15 06:10:46 +0000 UTC
#235     Centril Cross linking: Document toolstate rules 2019-07-05 08:27:32 +0000 UTC
#62861 matthiask submodules: update clippy from 164310dd to 49ff0d9d 2019-07-22 01:16:33 +0000 UTC
#62804 lundibund rustc_typeck: improve diagnostics for _ const/static de 2019-07-19 16:01:03 +0000 UTC
#64      Centril Account for rest patterns (`..`) 2019-07-10 23:31:56 +0000 UTC
#1465    lu-zero Make ilog64 const 2019-07-20 17:09:28 +0000 UTC
#382   todo[bot] Missing fminf in compiler-builtins for soft-float 2019-07-20 11:46:52 +0000 UTC
#1334   ppannuto Tracking: Deny warnings for documentation builds 2019-07-11 22:51:34 +0000 UTC
#60966   oli-obk Add a "diagnostic item" scheme for lints referring to l 2019-05-19 18:17:26 +0000 UTC
#62848   matklad Use unicode-xid crate instead of libcore 2019-07-21 12:51:16 +0000 UTC
#826   luckypoem on mac,"cargo build --release" encounter errors 2019-07-15 16:39:59 +0000 UTC
#1466      csmoe async unsafe fn syntax 2019-06-30 13:48:31 +0000 UTC
#7071   yaahallo Add experimental feature to link against local copy of  2019-06-26 00:31:24 +0000 UTC
#172   carols10c Tracking issue for more slice patterns 2019-06-14 19:36:38 +0000 UTC
#46        ehuss Add unstable attributes on function arguments 2019-06-28 02:50:58 +0000 UTC
#43      Centril Implement the bors auto r?eassignment on r+ policy 2019-05-30 21:28:17 +0000 UTC
#62262    varkor Extend `#[must_use]` to nested structures 2019-07-01 01:03:47 +0000 UTC
#49      Centril Prevent rollups in the web UI from including rollup=nev 2019-06-22 07:28:47 +0000 UTC
#62108      Zoxc Use sharded maps for queries 2019-06-25 02:34:56 +0000 UTC
47295 issues:
#62134   Xanewok Accept multiple --error-format flags 2019-06-25 23:27:08 +0000 UTC
#49        ehuss Update suffixed literal support 2019-06-28 02:52:27 +0000 UTC
#44        ehuss Add unstable new await syntax 2019-06-28 02:50:10 +0000 UTC
#2     non-binar Gameplan 2019-07-11 16:23:49 +0000 UTC
#1499    Xanewok Investigate switching to Azure Pipelines for Windows 2019-06-26 20:27:40 +0000 UTC
#5     jean-airo Forced serde_derive dependency 2019-07-19 23:24:05 +0000 UTC
#62225  RalfJung Spurious RLS test failures 2019-06-29 08:57:23 +0000 UTC
#3672  rossmacar Handle double semicolon 2019-07-06 17:43:42 +0000 UTC
#173   carols10c Tracking issue for union improvements 2019-06-14 19:59:10 +0000 UTC
#955   mikemorri stop auto-wrapping code snippets in fn main() 2019-06-14 17:36:46 +0000 UTC
#62839    Kobzol DWARF variant metadata for compressed enums with niche  2019-07-20 22:15:25 +0000 UTC
#170   carols10c Tracking issue for async/await 2019-06-13 23:12:52 +0000 UTC
#62846  RalfJung tidy complains about valid stability attribute 2019-07-21 11:01:45 +0000 UTC
#7           anp Enable clippy in CI 2019-06-21 06:18:49 +0000 UTC
#47    parsley42 Better docs / rendering 2019-07-02 19:04:39 +0000 UTC
#62711    lzutao Which is accessor method of MetadataExt? 2019-07-16 03:54:03 +0000 UTC
#62566 nikomatsa include material on async fn in the Rust reference 2019-07-10 18:41:32 +0000 UTC
#62207   parched Implement va_arg for AArch64 Linux 2019-06-28 14:04:49 +0000 UTC
#185   rusty-sna errors3.rs has outdated informations 2019-07-01 06:51:08 +0000 UTC
#1046  Thomasdez Determine how to do deal with WSAStartup on Windows 2019-07-22 10:09:45 +0000 UTC
#13      AlexEne Gamedev section for the rust-lang.org website. 2019-06-26 20:07:55 +0000 UTC
#150    RalfJung Collection of assumptions about MIR semantics 2019-06-23 12:11:56 +0000 UTC
#252   dependabo Dependabot can't resolve your Rust dependency files 2019-07-09 08:16:44 +0000 UTC
#625   petrochen Document that expressions in non-trailing expression st 2019-06-20 21:04:17 +0000 UTC
#61845      Zoxc Use a sharded dep node to dep node index map 2019-06-14 18:39:54 +0000 UTC
#5        oxr463 Continuous Integration 2019-06-21 12:54:15 +0000 UTC
#1          drom References 2019-06-26 21:33:02 +0000 UTC
#1675  alexcrich Update 'threads-xform' for LLVM 9 2019-07-19 18:11:54 +0000 UTC
#4     yoshuawuy key-value logging 2019-07-21 14:18:03 +0000 UTC
#300      gnzlbg Migrate libc to Azure Pipelines 2019-07-07 10:20:00 +0000 UTC
47295 issues:
#300      gnzlbg Migrate libc to Azure Pipelines 2019-07-07 10:20:00 +0000 UTC
#904   inscroggn resize image 2019-07-12 16:59:16 +0000 UTC
#430     Schaeff Track deprecated warning in lazy_static 2019-07-09 15:29:34 +0000 UTC
#168   carols10c Tracking issue for exception about `extern crate` with  2019-06-13 22:11:49 +0000 UTC
#3649    Xanewok `FileLines::all().to_json_spans()` returns `[]` which i 2019-06-24 13:38:17 +0000 UTC
#223      varkor rust-highfive only warns about updated submodules when  2019-07-01 01:07:37 +0000 UTC
#62802 rust-high `miri` no longer builds after rust-lang/rust#62679 2019-07-19 14:11:44 +0000 UTC
#62803 rust-high `rls` no longer builds after rust-lang/rust#62679 2019-07-19 14:11:45 +0000 UTC
#46    EvertonMe curl -O my-first-transaction.md #fix requeriments   2019-07-08 17:54:27 +0000 UTC
#901     dtolnay Wrong Rust logo in social media images 2019-07-11 00:52:04 +0000 UTC
#22     Arnavion Update for error-chain v0.12.0 2018-06-30 23:19:12 +0000 UTC
#145    RalfJung Validity of Box<T> 2019-06-18 15:35:48 +0000 UTC
#1977  kevinmeha Replace deprecated `...` range syntax with `..=` 2019-06-04 05:20:54 +0000 UTC
#13          g-k [rust] link to, incorporate, or track other rust guides 2019-05-03 17:38:21 +0000 UTC
#125       eddyb Improve printing of expected patterns in parse errors. 2019-07-18 20:02:23 +0000 UTC
#46      Centril Check for "No output has been received in the last 30m0 2019-06-07 22:21:28 +0000 UTC
#60559      Zoxc [WIP] Prerequisites from dep graph refactoring  2019-05-05 12:25:09 +0000 UTC
#6     That3Perc acquire_uninitialized should return MaybeUninit 2019-06-16 18:29:17 +0000 UTC
#4139   emmericp [WIP] RUN: enable --show-output when executing tests 2019-07-12 22:03:39 +0000 UTC
#22      oli-obk Dotenv codegen hangs forever 2019-06-27 13:30:59 +0000 UTC
#46    TylerReid Rewrite Rust in Rust 2019-06-07 17:16:16 +0000 UTC
#61693 matthiask no issue created when clippy toolstate breaks 2019-06-09 13:18:53 +0000 UTC
#799   Manishear [meta] Tracking issue for internationalization 2019-05-23 16:58:52 +0000 UTC
#49     josephlr Stop using dummy implementation 2019-06-29 02:30:35 +0000 UTC
#61976      jsgf Add Mutex::with 2019-06-19 22:25:05 +0000 UTC
#8        twilco Explore mdbooks for RISC-V from scratch series 2019-06-12 15:15:56 +0000 UTC
#607   LifeIsStr Use rust fix for vscode quickfixes ? 2019-06-20 12:27:33 +0000 UTC
#62715 rust-high `miri` no longer builds after rust-lang/rust#62704 2019-07-16 08:39:41 +0000 UTC
#846      oxr463 Pre-commit githook for enforcing coding style 2019-07-19 12:37:19 +0000 UTC
#2      tomByrer suggestion: + TOML 2019-06-17 22:46:21 +0000 UTC
47295 issues:
#572       ehuss `use` does not work with built-in macros. 2019-04-21 17:13:29 +0000 UTC
#846      oxr463 Pre-commit githook for enforcing coding style 2019-07-19 12:37:19 +0000 UTC
#174   carols10c Tracking issue for `std::simd` 2019-06-14 20:06:38 +0000 UTC
#61    rigelroza add https://github.com/rust-lang/rust.vim to install 2018-02-11 18:21:22 +0000 UTC
#637   nikomatsa document that impls may be more general than traits whe 2019-07-11 22:37:36 +0000 UTC
#4182  alexander let_and_return "known problems" 2019-06-07 15:13:34 +0000 UTC
#61879  stjepang Stabilize todo macro 2019-06-15 23:26:29 +0000 UTC
#61366    lzutao Stabilize exact_size_is_empty feature 2019-05-30 18:38:22 +0000 UTC
#237   nico-abra Replace static mut bool with AtomicBool 2019-07-15 05:58:28 +0000 UTC
#1565    lnicola CI: ra_vfs and proptest are not cached 2019-07-20 13:37:03 +0000 UTC
#70    thejpster Handle hex numbers 2019-07-16 15:41:12 +0000 UTC
#99    brettcann Add rustfmt and clippy to Rust dev container 2019-07-09 18:59:50 +0000 UTC
#7143  alexcrich Enable pipelined compilation by default 2019-07-17 21:12:38 +0000 UTC
#62522 SimonSapi Is Rc’s and Arc’s data_offset correct? 2019-07-09 11:22:44 +0000 UTC
#25    ashthespy  Bump nix to fix arm compilation issues caused by `PTRA 2019-07-15 21:41:41 +0000 UTC
#62254   Centril Tracking issue for `#![feature(slice_patterns)]` 2019-06-30 21:07:53 +0000 UTC
#77       gnzlbg Segfault while building the Rust toolchain with a toolc 2019-07-04 17:27:41 +0000 UTC
#178   carols10c Tracking issue for wasm-unknown-unknown 2019-06-14 21:12:48 +0000 UTC
#76      Centril Be forward compatible with rust-lang/rust#59928 2019-04-30 11:26:52 +0000 UTC
#62755    alexwl Type inference in the presence of recursive impls may r 2019-07-17 15:30:17 +0000 UTC
#62022   oli-obk Deduplicate error messages 2019-06-21 10:55:46 +0000 UTC
#1752    Nemo157 Should BoxFuture/BoxStream require Sync? 2019-07-20 09:30:23 +0000 UTC
#175   carols10c Tracking issue for installing rust without docs via rus 2019-06-14 20:16:26 +0000 UTC
#13    mathewcoh Update README to suggest latest version of crate 2019-07-13 06:32:37 +0000 UTC
#169   carols10c Tracking issue for `?` on either Option or Result in th 2019-06-13 22:24:44 +0000 UTC
#2604  stevenson rustfmt error overwrites buffer 2019-06-19 15:29:21 +0000 UTC
#59404      Zoxc Make ongoing_codegen a query 2019-03-24 21:06:10 +0000 UTC
#638   khwerhahn Cargo install --path jormungandr fails 2019-07-16 07:22:31 +0000 UTC
#336         djc Code coverage is broken 2019-04-28 19:47:58 +0000 UTC
#59338      Zoxc Turn parsing into a query 2019-03-21 10:42:17 +0000 UTC
47295 issues:
#59282      Zoxc Turn macro expansion and name resolution into a query 2019-03-18 19:54:35 +0000 UTC
#113   jakeswens fixing for rust 1.36.0 2019-07-09 22:04:44 +0000 UTC
#5      kilpatty Convert ok_or to ? once try_trait is implemented 2019-06-18 16:45:19 +0000 UTC
#22    nikomatsa "surprisingly hard things" 2019-07-09 17:25:40 +0000 UTC
#60938 jonas-sch rustdoc: make #[doc(include)] relative to the containin 2019-05-18 14:14:45 +0000 UTC
#62086 petrochen Define built-in macros through libcore 2019-06-23 20:08:47 +0000 UTC
#62144     yshui Explanation of E0207 is not what actually happens 2019-06-26 09:12:08 +0000 UTC
#425     Schaeff [WIP] Curve-generic ZoKrates 2019-07-08 14:22:33 +0000 UTC
#7088    matklad Forbid setting `RUSTC_BOOTSTRAP` from a build script 2019-07-03 06:42:13 +0000 UTC
#891        BO41 bindgen max_align_t workaround can be fixed 2019-05-08 15:56:12 +0000 UTC
#62280  cramertj Tracking issue for `slice_take` 2019-07-01 21:02:27 +0000 UTC
#62111   Centril Add a method for computing the absolute difference betw 2019-06-25 08:05:44 +0000 UTC
#23558       jdm Upgrade servo_rand to newer rand releases 2019-06-12 18:56:35 +0000 UTC
#61821  estebank 3 duplicated errors on index out of bound const evaluat 2019-06-14 01:26:47 +0000 UTC
#59904      Zoxc Remove queries from rustc_interface 2019-04-12 04:45:50 +0000 UTC
#568       ehuss Document name resolution. 2019-04-21 17:10:15 +0000 UTC
#270          64 Support zeroed allocations (e.g `reserve_zeroed`, `with 2019-07-07 17:34:55 +0000 UTC
#14    Thomasdez [WIP] Print key values 2019-04-25 23:22:23 +0000 UTC
#7           kpp Implement stream combinators using async_stream_block a 2019-07-01 08:18:05 +0000 UTC
#616   nicodemus 1.36 borrow checker warns about possible undefined beha 2019-07-05 22:22:43 +0000 UTC
#1726    ts25504 [MAINTENANCE] Use rust-toolchain to specify a rustc ver 2019-06-17 09:53:22 +0000 UTC
#39      coord-e GitBook is deprecated 2019-06-29 08:26:28 +0000 UTC
#61415    varkor Use const generics for array impls 2019-05-31 21:38:19 +0000 UTC
#62177   Centril Move some tests in src/test/compile-fail -> src/test/ui 2019-06-27 12:20:12 +0000 UTC
#283   calixtema Use rust hashmap instead of fxhashmap 2019-04-27 12:45:46 +0000 UTC
#140    lopopolo Upgrade bindgen to 0.50.0 2019-07-03 07:07:06 +0000 UTC
#636   nikomatsa document when the values of default type parameters are 2019-07-11 22:33:17 +0000 UTC
#45        ehuss Add unstable C variadic function argument 2019-06-28 02:50:38 +0000 UTC
#42        ehuss Update for if_while_or_patterns 2019-06-28 02:49:20 +0000 UTC
#998         ia0 Fix meta variable misuse 2019-07-17 21:42:45 +0000 UTC
47295 issues:
#8     ChrisLinn use `fold` for single_number Solution1 2019-07-19 00:11:37 +0000 UTC
#4194    Centril New lint: for `#![feature(associated_type_bounds)]` (Ex 2019-06-11 12:08:04 +0000 UTC
#845   BartMasse luhn-trait description has dead link to wrong informati 2019-07-21 22:50:18 +0000 UTC
#14    todo[bot] use clamp() when stabilized 2019-06-19 03:02:48 +0000 UTC
#4        htfy96 Use atomic_element_unordered_copy_memory_nonoverlapping 2019-06-02 23:47:09 +0000 UTC
#59205      Zoxc Turn HIR lowering into a query 2019-03-15 10:38:56 +0000 UTC
#2        iceiix Implement std::io::Read on &File, in addition to File 2019-05-29 15:34:51 +0000 UTC
#786       skade Fix outstanding bugs in the /learn page 2019-05-21 13:50:07 +0000 UTC
#4287  matthiask rustup https://github.com/rust-lang/rust/pull/62679/ 2019-07-19 14:41:58 +0000 UTC
#772       skade Rework "Rust in Production" section 2019-05-15 09:30:58 +0000 UTC
#42327  scottmcm Tracking issue for `ops::Try` (`try_trait` feature) 2017-05-31 08:23:55 +0000 UTC
#1513    Xanewok Collect file -> edition mapping after AST expansion 2019-07-14 20:15:24 +0000 UTC
#4083  Manishear Audit macro checks and see if they can ignore desugarin 2019-05-12 03:45:52 +0000 UTC
#61997 nikomatsa Tracking issue for member constraints in region inferen 2019-06-20 14:00:36 +0000 UTC
#3837  Manishear Move cargo-clippy into cargo 2019-03-03 06:27:45 +0000 UTC
#2         c-edw ValueError when $SHELL is not an absolute path. 2019-06-17 17:35:06 +0000 UTC
#4285  matthiask rustup https://github.com/rust-lang/rust/pull/62764 2019-07-18 22:36:58 +0000 UTC
#62496 gyakovlev install.sh miplaces codegen-backends directory with lib 2019-07-08 16:37:08 +0000 UTC
#70       tomtau Add Cargo.lock ? 2019-06-14 01:04:23 +0000 UTC
#6       Centril Be forward compatible with rust-lang/rust#59928 2019-04-30 12:09:51 +0000 UTC
#62678  flip1995 Implement VEC_NEW internal lint 2019-07-14 15:58:59 +0000 UTC
#14    ubnt-intr Set cfg(test) flag to generated item(s) 2019-06-25 07:27:30 +0000 UTC
#977     MOZGIII Support for `peer_addr` for `UdpSocket` 2019-06-12 18:42:22 +0000 UTC
#1     rasendubi Fix assert! macro usage 2019-04-25 08:03:08 +0000 UTC
#31    todo[bot] CanonicalSignal should be a trait alias once stabilized 2019-06-24 04:56:27 +0000 UTC
#1493   LegNeato Optionally use #[repr(transparent)] instead of type ali 2019-01-17 01:52:50 +0000 UTC
#385         kpp Replace IpPort::is_global with IpAddr::is_global from s 2019-06-01 11:47:03 +0000 UTC
#3835    phansch Provide an easy way for contributors to run cargo fmt l 2019-03-02 08:19:48 +0000 UTC
#1517  matthiask rustup https://github.com/rust-lang/rust/pull/62679/ 2019-07-19 19:20:15 +0000 UTC
#304        da-x `__rust_probestack` needs CFI annotations to assist in  2019-07-19 14:46:06 +0000 UTC
47295 issues:
#304        da-x `__rust_probestack` needs CFI annotations to assist in  2019-07-19 14:46:06 +0000 UTC
#62210 QuietMisd Tracking issue for `cfg(doctest)` 2019-06-28 15:09:12 +0000 UTC
#59302   mati865 Tracking issue for musl host toolchain 2019-03-19 23:10:34 +0000 UTC
#16    yingxiong Integer suffixes discarded 2019-07-19 17:46:45 +0000 UTC
#1     rasendubi Fix assert! macro usage 2019-04-25 07:56:41 +0000 UTC
#157   rasendubi Fix assert! macro usage 2019-04-25 07:13:42 +0000 UTC
#318     palfrey Allow failures on nightly because it's not always stabl 2019-07-03 15:56:26 +0000 UTC
#2     rasendubi Fix assert! macro usage 2019-04-25 07:52:10 +0000 UTC
#230   linkmauve Check rustc and cargo for Haiku 2019-06-18 11:43:32 +0000 UTC
#61810   phansch annotate-snippet emitter: Deal with multispans in macro 2019-06-13 19:27:51 +0000 UTC
#60547 jackpot51 redox: convert to target_family unix 2019-05-04 16:07:31 +0000 UTC
#62569 joshtripl Should Rust still ignore SIGPIPE by default? 2019-07-10 19:56:27 +0000 UTC
#141   tiziano88 New example: rustfmt 2019-07-11 22:41:46 +0000 UTC
#52    pietroalb Ignore commands in quotes 2019-06-29 14:13:15 +0000 UTC
#62370   Nemo157 Tracking issue for Box::into_pin (feature `box_into_pin 2019-07-04 10:54:44 +0000 UTC
#5           kpp Reimplement stream combinators via generators 2019-06-28 22:35:14 +0000 UTC
#53667   Centril Tracking issue for eRFC 2497, "if- and while-let-chains 2018-08-24 12:25:32 +0000 UTC
#1964    PunKeel The dangle function should be marked as "Does not compi 2019-05-20 13:51:22 +0000 UTC
#46     TheDan64 Tracking Issue: Translation nightly features 2019-04-02 19:23:03 +0000 UTC
#1432      eoger Use reqwest w/ rusttls on Linux 2019-07-19 18:28:04 +0000 UTC
#1282  edwin0che `CargoWorkspace` do not contains non-normal dependencie 2019-05-17 15:43:54 +0000 UTC
#62729  roblabla Link errors when compiling for i386 with +soft-float 2019-07-16 20:19:11 +0000 UTC
#1043   lgalabru Considering untangling interdependent modules 2019-07-02 17:55:21 +0000 UTC
#51        ehuss Should MacroCall include optional ident? 2019-06-28 02:53:18 +0000 UTC
#1403  jackpot51 Redox CI 2019-06-21 02:52:51 +0000 UTC
#57173      Zoxc Allocate HIR on an arena 2018-12-28 16:43:30 +0000 UTC
#194         jdm Integrate try/retry fix from rust's fork 2019-05-30 17:31:25 +0000 UTC
#129    jethrogb Support cross-compiling doctests 2019-04-17 00:49:57 +0000 UTC
#291   ishitatsu Proposal: migrate to Azure pipelines 2019-06-26 11:32:37 +0000 UTC
#62534       jdm Add a crt-debug target feature 2019-07-09 15:50:14 +0000 UTC
47295 issues:
#11      Centril Be forward compatible with rust-lang/rust#59928 2019-04-30 12:30:05 +0000 UTC
#618      matiit Successfully installed notification doesn't display pro 2019-07-01 14:37:21 +0000 UTC
#177   carols10c Tracking issue for llvm-tools-preview becoming not prev 2019-06-14 20:25:25 +0000 UTC
#171   carols10c Tracking issue for using impl Trait in more places 2019-06-14 19:29:47 +0000 UTC
#630    zygoloid Rust reference is missing a description of the visibili 2019-07-02 23:21:28 +0000 UTC
#16       bjorn3 Support triples of form `x86_64-apple-macosx10.7.0` 2019-05-15 16:52:24 +0000 UTC
#36       ytmhub Compiler panics on 'cargo doc' 2019-06-21 17:21:06 +0000 UTC
#165    nlhepler refactor div-conq SVD, impl svd_rand 2019-06-27 23:19:18 +0000 UTC
#7017  MageSlaye TypeId is different between compilations 2019-06-06 15:05:58 +0000 UTC
#1741  Thomasdez Add AsyncWriteExt::write_all_vectored utility 2019-07-16 16:26:52 +0000 UTC
#1     edwin0che Some notes about the panic 2019-05-26 18:34:41 +0000 UTC
#2           nrc Use the crates.io version of libsyntax instead of the s 2017-05-08 01:39:02 +0000 UTC
#296   ishitatsu Update compiler for the Android builder 2019-07-03 12:12:19 +0000 UTC
#62293   Centril Unsupport the `await!(future)` macro 2019-07-02 04:51:29 +0000 UTC
#23    steveklab Add the complete set of rules for method resolution in  2017-03-23 15:35:11 +0000 UTC
#1     rasendubi Fix assert! macro usage 2019-04-25 07:49:53 +0000 UTC
#68     phil-opp Only include dependencies when `binary` feature is enab 2019-07-18 09:47:09 +0000 UTC
#31    rollsmorr Broken MIR: generator contains type  ... 2019-06-18 13:19:51 +0000 UTC
#176     phansch Use resume_unwind instead of panic!() for nicer errors 2019-04-16 17:15:43 +0000 UTC
#1        tinaun bot consolidation + oxidation 2019-06-13 02:48:46 +0000 UTC
#62633   cuviper Tracking issue for Option::expect_none(msg) and unwrap_ 2019-07-12 20:07:13 +0000 UTC
#7         denzp Use `libLLVM` instead of `librustc_codegen_llvm` 2019-04-25 14:49:47 +0000 UTC
#9     contradic Unrecognized option: 'write-mode' when applying the rus 2019-07-15 09:13:05 +0000 UTC
#63653   nmattia Cargo build without cargoSha256, use Cargo.lock instead 2019-06-22 14:14:28 +0000 UTC
#62799  RalfJung use const array repeat expressions for uninit_array 2019-07-19 12:48:16 +0000 UTC
#62     jrmuizel Support using rust-semverver to detect appropriate vers 2018-04-18 19:54:39 +0000 UTC
#1369  BaoshanPa How to use my local libc when building rust from souce  2019-05-25 18:37:13 +0000 UTC
#43466       nrc Tracking issue for RFC 1946 - intra-rustdoc links 2017-07-25 03:35:26 +0000 UTC
#60406   Centril Tracking issue for RFC 2565, "Attributes in formal func 2019-04-30 10:09:24 +0000 UTC
#12     appaquet Upstream capnp thread safe modification once we have At 2019-02-13 02:46:49 +0000 UTC
分類結果
1ヶ月未満 191
1年未満 102
1年以上 7
//!-textoutput
*/
