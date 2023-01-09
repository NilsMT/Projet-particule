
## initialisation
cd existing_folder

git init --initial-branch=main

git remote add origin https://gitlab.univ-nantes.fr/E221936K/systeme-de-particule.git

git add .

git commit -m "Initial commit"

git push -u origin main

## Pour update  (à l'IUT)

git add . 

*selectionnne les fichiers*

git commit -m "mise à jour"

*regroupe dans une archive (.zip)*

git push 

*l'envoie au git*                                           

## Pour le récupérer (initialisation)

git clone https://gitlab.univ-nantes.fr/E221936K/systeme-de-particule.git

*récupère le répertoire*

## Pour le récupérer (tout le temps)

git pull

*récupère le répertoire initialisé et le met à jour automatiquement*
