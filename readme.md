# Example for define templates for emails

Most server-side languages have techniques for taking templates. Go has adopted a relatively simple scripting language in the template package. This is small example of how to use it as email templates

## Templates definition
We have example for sending email to customer and shop owner when new order is created.

```
tpls
    └── emails
        ├── layout.html
        ├── orderItems.html
        ├── orderNewAdmin.html
        └── orderNewCustomer.html
```

## TplParser


```
type TplParser struct {
	Tpls map[string]*template.Template
}

func NewTplParser(tplNames []string) *TplParser {
	es := &TplParser{Tpls: make(map[string]*template.Template, 0)}

	for _, tplName := range tplNames {
		t, err := template.ParseFiles(
			"./tpls/emails/layout.html",
			"./tpls/emails/orderItems.html",
			"./tpls/emails/"+tplName+".html",
		)
		if err != nil {
			fmt.Println("Error in loading email templates")
			panic(err)
		}
		es.Tpls[tplName] = t
	}
	return es
}

func (es *TplParser) Parse(emailType string, data Order) (string, error) {
	t := es.Tpls[emailType]
	var tplBytes bytes.Buffer

	if err := t.ExecuteTemplate(&tplBytes, "layout", data); err != nil {
		return "", err
	}

	return tplBytes.String(), nil
}

```

Parsing email template for `orderNewCustomer`

```
tp := NewTplParser([]string{
    "orderNewCustomer",
    "orderNewAdmin",
})

order := Order{
    Date:    time.Now(),
    Billing: Address{"Rabeeshkumar", "A R"},
    Number:  "12320",
    Items:   []Item{Item{"P1", 2, 23.23}, Item{"P2", 8, 99.01}},
}

// get email template for customer
customerEmailContent, err := tp.Parse("orderNewCustomer", order)
if err != nil {
    panic(err)
}
fmt.Println(customerEmailContent)
```

## Template definition

Layout template
```
{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>Exmaple email template</title>
</head>
<body>
  <div>
    {{ template "content" . }}
  </div>
</body>
</html>
{{ end }}
```

Admin order create email template
```
{{ define "content" }}
<p>
    You have received an order from {{ .Billing.Firstname }} {{ .Billing.LastName }}.
    The order is as follows
</p>
{{ template "orderItems" . }}
{{ end }}
```

Customer order create email template
```
{{ define "content" }}
<p>Dear <strong>{{ .Billing.Firstname }} {{ .Billing.LastName }}</strong>,</p>

<p>
  Thank you for your order - this is your order confirmation.
  <br />
  Order date	{{ .Date.Format "Jan_2 Mon 2006 15:04:05" }}
  <br />
  Order	{{ .Number }}
</p>

{{ template "orderItems" . }}
{{ end }}
```

Order items listing template
```
{{ define "orderItems" }}
<div>
  <table border="0">
    <thead>
      <tr >
        <th class="td" scope="col">Product</th>
        <th class="td" scope="col">Quantity</th>
        <th class="td" scope="col">Price</th>
      </tr>
    </thead>
    <tbody>
      {{range $idx, $item := .Items}}
      <tr>
        <td class="td" scope="col">{{$item.Name}}</td>
        <td class="td" scope="col">{{$item.Quantity}}</td>
        <td class="td" scope="col">{{$item.Total}} INR</td>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>
{{ end }}
```

## MIT
