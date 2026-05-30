# Health Checks

### Readiness vs Liveness Probe

#Readiness
Prüft ob der Pod bereit ist Traffic zu emmpfangen. Wenn es fehlschlägt wird der Pod aus dem Load Balancer
entfernt. Ist nützlich beim Start da die App Zeit zum hochfahren bracuht

#Liveness
Prüft ob der Pod noch funktioniert, wenn es fehlschlägt wir der Pod neugestartet

### Was passiert wenn eine Probe fehlschlägt?
- Readiness: Pod bekommt keinen Traffic, also keinen Neustart
- Liveness: Pod wird von Kubernetes automatisch neu gestartet

## Resource Limits

### Was passiert wenn Memory/CPU Limit überschritten wird?
- Memory: Pod wird sofort beendet und neu gestartet
- CPU: Pod wird gedrosselt, läuft langsamer aber
  nicht beendet

### Warum requests und limits angeben?
- Requests: Kubernetes reserviert diese Ressourcen für den Pod
  beim Scheduling, garantierte Mindestressourcen
- Limits: Verhindert dass ein Pod zu viele Ressourcen verbraucht
  und andere Pods beeinträchtigt
- Ohne requests: Kubernetes kann Pod auf überlasteten Node
  schedulen
- Ohne limits: Ein Pod kann alle Ressourcen eines Nodes aufbrauchen

