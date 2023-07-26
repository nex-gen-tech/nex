package router

import (
	"regexp"
	"strings"
)

const (
	simpleParam paramType = iota
	wildcardParam
	regexParam
)

var regexCache = make(map[string]*regexp.Regexp)

type paramType int

type paramMatcher struct {
	ptype paramType
	regex *regexp.Regexp
	name  string
}

func newParamMatcher(segment string) (*paramMatcher, error) {
	pm := &paramMatcher{}

	if segment == "*" {
		pm.ptype = wildcardParam
	} else if segment[0] == ':' {
		idxStart, idxEnd := strings.Index(segment, "("), strings.LastIndex(segment, ")")
		if idxStart != -1 && idxEnd != -1 {
			pm.ptype = regexParam
			pm.name = segment[1:idxStart]
			regexPattern := segment[idxStart : idxEnd+1]
			if cachedRegex, exists := regexCache[regexPattern]; exists {
				pm.regex = cachedRegex
			} else {
				var err error
				pm.regex, err = regexp.Compile("^" + regexPattern + "$")
				if err != nil {
					return nil, err
				}
				regexCache[regexPattern] = pm.regex
			}
		} else {
			pm.ptype = simpleParam
			pm.name = segment[1:]
		}
	} else {
		return nil, nil
	}

	return pm, nil
}

func (pm *paramMatcher) match(segment string) (bool, string) {
	switch pm.ptype {
	case simpleParam, wildcardParam:
		return true, segment
	case regexParam:
		if pm.regex.MatchString(segment) {
			return true, segment
		}
	}
	return false, ""
}
