package gatherer

// Gatherer should implement ...
type Gatherer interface {
	Set(...interface{})
	Run()
	Report() map[string]interface{}
}

// Init Gatherers.
func Init() map[string]Gatherer {
	return map[string]Gatherer{
		"basicInfo": NewBasicInfo(),
		"bypassCF":  NewCFBypass(),
		"whois":     NewWhois(),
		"cms":       NewCMSDetector(),
		"honeypot":  NewHoneypotDetecter(),
		"port":      NewPortScanner(),
		"tracert":   NewTracer(),
		"subdomain": NewSubDomainScan(),
		"dirb":      NewDirBruter(),
	}
}
