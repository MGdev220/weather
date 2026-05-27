# Atelier Météo — Parsing JSON / XML en Go

## Partie A — Comparaison des schémas de données

Avant de modéliser les données en Go, voici l'analyse des deux formats sources :

| Donnée                   | Comment c'est représenté en JSON ?                                                      | Comment c'est représenté en XML ?                                                                             |
|:-------------------------|:----------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------------------------|
| **Pays**                 | Champ `country` en String à la racine de la station.                                    | Code ISO à 2 lettres dans l'attribut `country` de la station.                                                 |
| **Coordonnées**          | Objet imbriqué `location` contenant les champs `longitude` et `latitude` en float.      | Attributs `lat` et `lon` sur l'élément enfant `<coordinates>` dans l'element parent Station.                  |
| **Altitude**             | Champ `altitude_m` en int.                                                              | Attribut `altitude` situé sur l'élément enfant `<coordinates>`dans l'element parent Station.                  |
| **Modèle de capteur**    | Objet imbriqué `device` avec `type`, `manufacturer` et `installed_on` (tous en String). | Attributs `vendor`, `model`, et `since` sur l'élément enfant `<hardware>`, parent station.                    |
| **Température**          | Champ `temperature_celsius` en float dans l'observation.                                | Contenu textuel (chardata) d'un élément `<measure>` possédant l'attribut `type="temperature"`.                |
| **Conditions ciel**      | Champ `conditions` en String dans l'observation.                                        | Attribut `sky` placé directement sur l'élément `<observation>` parent station.                                |
| **Vent**                 | Objet `wind` imbriqué avec `speed_kmh` (float) et `direction_deg` (int).                | Attributs `speed` et `direction` sur l'élément enfant `<wind>` parent station.                                |
| **Notes (optionnelles)** | Champ `notes` en String dans l'observation (peut être `null`).                          | Élément `<note>` qui contient le texte, mais qui est totalement omis s'il n'y a pas de note. Dasn observation |

## Description des fichiers principal qui charge les deux sources, vérifie leur cohérence et affiche les résultats des requêtes.