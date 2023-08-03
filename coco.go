package main

type tag uint8
const (
  // special tags
  ufo tag = iota
	eof
  doc

  // literals
	name
	num
	str
	str_inter
	chr

  // operators
  mut
	add
	sub
	mul
	quo
	rem
	aeq
	seq
	meq
	qeq
	req
	lt
	le
	gt
	ge
	eq
	ne
	and
	or
	xor
	nand
	nor
	shl
	shr
	flip
	not

  // delimeters
	dot  // .
	rng  // ..
	com  // ,
	col  // :
	sem  // ;
	opt  // ?
	aro  // ->
	at   // @
	as   // #
	lcu  // {
	rcu  // }
	lpn  // (
	rpn  // )
	lsq  // [
	rsq  // ]
  
  // keywords
	let_kw
	var_kw
	if_kw
	else_kw
	when_kw
	then_kw
	for_kw
	in_kw
  notin_kw
	stop_kw
	skip_kw
	enum_kw
	data_kw
	type_kw
	obj_kw
	fn_kw
	ret_kw
	go_kw
	pub_kw
	use_kw
	mod_kw
	end_kw
	and_kw
	or_kw
	is_kw
	isnot_kw
	yes_kw
	no_kw
)

var keywords = map[string] tag {
  "let": let_kw,
  "var": var_kw,
  "if": if_kw,
  "else": else_kw,
  "when": when_kw,
  "then": then_kw,
  "for": for_kw,
  "in": in_kw,
  "!in": notin_kw,
  "stop": stop_kw,
  "skip": skip_kw,
  "enum": enum_kw,
  "data": data_kw,
  "type": type_kw,
  "obj": obj_kw,
  "fn": fn_kw,
  "ret": ret_kw,
  "go": go_kw,
  "pub": pub_kw,
  "use": use_kw,
  "mod": mod_kw,
  "end": end_kw,
  "and": and_kw,
  "or": or_kw,
  "is": is_kw,
  "is!": isnot_kw,
  "yes": yes_kw,
  "no": no_kw,
}
