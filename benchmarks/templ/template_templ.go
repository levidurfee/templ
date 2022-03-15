// Code generated by templ@v0.2.185 DO NOT EDIT.

package testhtml

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Render(p Person) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<h1>")
		if err != nil {
			return err
		}
		// StringExpression
		_, err = templBuffer.WriteString(templ.EscapeString(p.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" style=\"font-family: &#39;sans-serif&#39;\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" id=\"test\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" data-contents=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(`something with "quotes" and a <tag>`))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div>")
		if err != nil {
			return err
		}
		// Text
		var_2 := `email:`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<a")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" href=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		var var_3 templ.SafeURL = templ.URL("mailto: " + p.Email)
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_3)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// StringExpression
		_, err = templBuffer.WriteString(templ.EscapeString(p.Email))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<hr")
		if err != nil {
			return err
		}
		// Element Attributes
		if true {
			_, err = templBuffer.WriteString(" noshade")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<hr")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" optionA")
		if err != nil {
			return err
		}
		if true {
			_, err = templBuffer.WriteString(" optionB")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" optionC=\"other\"")
		if err != nil {
			return err
		}
		if false {
			_, err = templBuffer.WriteString(" optionD")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<hr")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" noshade")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

