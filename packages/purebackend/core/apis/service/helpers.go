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
		<!DOCTYPE html>
		<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
			<title>` + title + `</title>
			<link rel="preconnect" href="https://fonts.googleapis.com/">
			<link rel="preconnect" href="https://fonts.gstatic.com/" crossorigin="">
			<style type="text/css">
			.button:active {
				border-style: outset;
			}
			</style>
		<style data-jss="" data-meta="MuiDialog">
		@media print {
		.MuiDialog-root {
			position: absolute !important;
		}
		}
		.MuiDialog-scrollPaper {
		display: flex;
		align-items: center;
		justify-content: center;
		}
		.MuiDialog-scrollBody {
		overflow-x: hidden;
		overflow-y: auto;
		text-align: center;
		}
		.MuiDialog-scrollBody:after {
		width: 0;
		height: 100%;
		content: "";
		display: inline-block;
		vertical-align: middle;
		}
		.MuiDialog-container {
		height: 100%;
		outline: 0;
		}
		@media print {
		.MuiDialog-container {
			height: auto;
		}
		}
		.MuiDialog-paper {
		margin: 32px;
		position: relative;
		overflow-y: auto;
		}
		@media print {
		.MuiDialog-paper {
			box-shadow: none;
			overflow-y: visible;
		}
		}
		.MuiDialog-paperScrollPaper {
		display: flex;
		max-height: calc(100% - 64px);
		flex-direction: column;
		}
		.MuiDialog-paperScrollBody {
		display: inline-block;
		text-align: left;
		vertical-align: middle;
		}
		.MuiDialog-paperWidthFalse {
		max-width: calc(100% - 64px);
		}
		.MuiDialog-paperWidthXs {
		max-width: 444px;
		}
		@media (max-width:507.95px) {
		.MuiDialog-paperWidthXs.MuiDialog-paperScrollBody {
			max-width: calc(100% - 64px);
		}
		}
		.MuiDialog-paperWidthSm {
		max-width: 600px;
		}
		@media (max-width:663.95px) {
		.MuiDialog-paperWidthSm.MuiDialog-paperScrollBody {
			max-width: calc(100% - 64px);
		}
		}
		.MuiDialog-paperWidthMd {
		max-width: 960px;
		}
		@media (max-width:1023.95px) {
		.MuiDialog-paperWidthMd.MuiDialog-paperScrollBody {
			max-width: calc(100% - 64px);
		}
		}
		.MuiDialog-paperWidthLg {
		max-width: 1280px;
		}
		@media (max-width:1343.95px) {
		.MuiDialog-paperWidthLg.MuiDialog-paperScrollBody {
			max-width: calc(100% - 64px);
		}
		}
		.MuiDialog-paperWidthXl {
		max-width: 1920px;
		}
		@media (max-width:1983.95px) {
		.MuiDialog-paperWidthXl.MuiDialog-paperScrollBody {
			max-width: calc(100% - 64px);
		}
		}
		.MuiDialog-paperFullWidth {
		width: calc(100% - 64px);
		}
		.MuiDialog-paperFullScreen {
		width: 100%;
		height: 100%;
		margin: 0;
		max-width: 100%;
		max-height: none;
		border-radius: 0;
		}
		.MuiDialog-paperFullScreen.MuiDialog-paperScrollBody {
		margin: 0;
		max-width: 100%;
		}
		</style></head>
		<body style="
			font-family: IBM Plex Sans;
			color: #475569;
			display: flex;
			justify-content: center;
			padding-top: 16px;
			">
			` + body + `
		</body>
	</html>
	`
}
