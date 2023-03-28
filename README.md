# Auto-connect-tool
Outil qui permet de  qui permet de changer automatiquement de réseau internet sur l'appareil  en fonction de la moins mauvaise performance du réseau.

- Ce code utilise la commande `mcli` pour récupérer la liste des interfaces réseau disponibles.
- Puis effectue un test de performance en utilisant la commande `ping` pour chaque interface réseau. 
- Il sélectionne ensuite automatiquement l'interface avec la meilleure performance et utilise à nouveau `nmcli` pour changer l'interface réseau active. 


