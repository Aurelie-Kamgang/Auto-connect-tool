package main
import (    
  "fmt"
  "os/exec"    
  "strings"
)
func main() {    
  // Récupération de la liste des interfaces réseau disponibles
    out, err := exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE,STATE", "device").Output()    
    if err != nil {
        fmt.Println("Erreur lors de la récupération de la liste des interfaces réseau :", err)        
        return
    }
  // Analyse de la sortie de nmcli pour obtenir la liste des interfaces actives    
     networkList := []string{}
     for _, line := range strings.Split(string(out), "\n") {       
       if line == "" {
            continue        
       }
       fields := strings.Split(line, ":")        
       if fields[1] == "wifi" || fields[1] == "ethernet" {
            if fields[2] == "connected" {                
            networkList = append(networkList, fields[0])
            }        
       }
    }
    // Test de la performance de chaque interface réseau    
    bestNetwork := ""
    bestPingTime := int(^uint(0) >> 1)    
    for _, network := range networkList {
        // Tester la performance en effectuant un ping vers google.com        
        pingCmd := exec.Command("ping", "-c", "3", "-q", "-w", "3", "google.com")
        pingCmd.Env = append(pingCmd.Env, fmt.Sprintf("DEVICE=%s", network))        
        if err := pingCmd.Run(); err != nil {
            fmt.Printf("Erreur lors du test de performance pour le réseau %s : %v\n", network, err)            
            continue
        }        
    pingOutput := string(pingCmd.Output())
        // Analyser la sortie de ping pour obtenir le temps de ping moyen
        pingTimeStr := strings.Split(strings.Split(pingOutput, "/")[4], ".")[0]        
        pingTime := 0
        fmt.Sscanf(pingTimeStr, "%d", &pingTime)
        // Mettre à jour le meilleur réseau si nécessaire        
       if pingTime < bestPingTime {
            bestNetwork = network            
            bestPingTime = pingTime
        }    
    }
    // Changer l'interface réseau active pour utiliser le meilleur réseau
    if bestNetwork != "" {        
      fmt.Printf("Changement de l'interface réseau active pour utiliser %s\n", bestNetwork)
      exec.Command("nmcli", "device", "connect", bestNetwork).Run()    
   }
}
