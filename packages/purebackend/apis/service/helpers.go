package service

import "net/mail"

func ValidateMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func BaseEmailTemplate(title, body string) string {
	return `
		<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
			<head>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta name="format-detection" content="telephone=no"/>
				<title>` + title + `</title>
				<style type="text/css">
					/* Resets: see reset.css for details */
					.ReadMsgBody { width: 100%; background-color: #ffffff;}
					.ExternalClass {width: 100%; background-color: #ffffff;}
					body {width: 100%; background-color: #ffffff; margin:0; padding:0; -webkit-font-smoothing: antialiased; font-family: Arial, Helvetica, sans-serif;}
					table {border-collapse: collapse;}
	
					@media only screen and (max-width: 640px)  {
						body[yahoo] .deviceWidth {width:440px!important; padding:0;}
						body[yahoo] .center {text-align: center!important;}
					}
					@media only screen and (max-width: 479px) {
						body[yahoo] .deviceWidth {width:280px!important; padding:0;}
						body[yahoo] .center {text-align: center!important;}
					}
				</style>
			</head>
			<body leftmargin="0" topmargin="0" marginwidth="0" marginheight="0" yahoo="fix" style="margin: 0px; padding: 0px; width: 100% !important; -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; -webkit-font-smoothing: antialiased; font-family: Arial, Helvetica, sans-serif;">
				` + body + `
			</body>
		</html>
	`
}