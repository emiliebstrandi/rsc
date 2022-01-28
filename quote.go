package quote 

import(
	rsc.io/quote

)

func Quote() string{
	return quote.Hello() + "\n" + quote.Go() + "\n" + quote.Glass() + "\n" + quote.Opt()
}