package models

import "encoding/json"

type Request struct {
	User        *UserDetails
	Org         *OrgDetails
	Body        []byte
	ParsedBody  map[string]interface{}
	Headers     map[string]string
	PathParams  map[string]string
	QueryParams map[string]string
}

type UserDetails struct {
	MailId string
}

type OrgDetails struct {
	Id string
}

func (r *Request) GetUserMail() string {
	return r.User.MailId
}

func (r *Request) GetOrgId() string {
	return r.Org.Id
}

func (r *Request) GetPathParam(param string) string {
	return getValueFromMap(r.PathParams, param)
}

func (r *Request) GetQueryParam(param string) string {
	return getValueFromMap(r.QueryParams, param)
}

func (r *Request) ParseJsonBody() {
	body := map[string]interface{}{}
	err := json.Unmarshal(r.Body, &body)
	if err != nil {
		panic(err)
	}
	r.ParsedBody = body
}

func (r *Request) GetParsedBodyAttribute(attributeName string) interface{} {
	if len(r.ParsedBody) == 0 {
		return nil
	}
	if val, ok := r.ParsedBody[attributeName]; ok {
		return val
	}
	return nil
}

func getValueFromMap(m map[string]string, param string) string {
	if val, ok := m[param]; ok {
		return val
	}
	return ""
}
