Projet réalisé par Tom Frémont et Nils Moreau--Thomas
## Sommaire
1. [Les fichiers](#les-fichiers)
2. [Panneau de configuration](#panneau-de-configuration)
3. [Liste des variables dynamique (configurable dynamiquement)](#liste-des-variables-dynamique)
4. [Liste des variables non dynamique (non configurable dynamiquement)](#liste-des-variables-non-dynamique)
5. [Faire les tests](#faire-les-tests)
## Les fichiers
Le système de particule se situe dans `ParticleSystem/src`, il contient :
- Le fichier `draw.go` qui gère l'affichage des particules et des vecteurs de direction
- Le fichier `particle/update.go` qui gère la modification des propriétés des particules
- Le fichier `config.json` qui contient la configuration que le système utilisera
- Le fichier `particle/extension_test.go` qui contient tout les test des extensions après la partie 4
- Le fichier `particle/rel1_test.go` qui contient tout les test jusqu'à la partie 4
- Le fichier `particle/extension.go` qui ccontient les fonctions relatives aux extensions
- Le fichier `particle/new.go` qui gère l'apparition des particules à l'initialisation
- Le fichier `particle/type` qui déclare les variable que contient une particule
- Le fichier `config/type.go` qui déclare ce que la `config.json` contient
## Panneau de configuration
La configuration est configurable dynamiquement (pendant l'exécution de la fenêtre) grâce à certaines touches
La méthode pour sélectionné se fait sur plusieurs niveau :
- il faut tout d'abord repérer le type de la variable que l'on veut modifier
- il faut sélectionner le type avec <kbd>A</kbd> et <kbd>E</kbd>
- il faut ensuite sélectionner la variable en question avec <kbd>←</kbd> et <kbd>→</kbd>
- il faut ensuite modifier la variable en question avec <kbd>↑</kbd> et <kbd>↓</kbd>

#### Action de chaque touche :
| Nom | Symbole | Action |
| ---- |---- | ---- |
| A | <kbd>A</kbd> | Avance dans la sélection des types |
| E | <kbd>E</kbd> | Recule dans la sélection des types |
| Flèche droite | <kbd>→</kbd> | Avance dans la sélection des variables |
| Flèche gauche | <kbd>←</kbd> | Avance dans la sélection des variables |
| Flèche haut | <kbd>↑</kbd> | Augmente la variable |
| Flèche bas | <kbd>↓</kbd> | Diminue la variable |

Il faut noté que en cas de sélection d'un booléen (bool) la touche qui augmente et qui diminue change juste la valeur à son opposé (si la variable est `true` elle devient `false` et inversement). 
Par ailleurs la sélection peut aussi de faire à l'envers : si l'ont recule en étant dans les bool, cela redirige aux sélection des formes, soit la dernière sélection), l'inverse est aussi possible (des formes aux bool)
## Liste des variables dynamique

#### Sélection numéro 1 (les bool)
| Type | Nom | Description |
| ---- |---- | ---- |
| bool | Debug |Si `true` : montre les FPS, le nombre de particule et les vecteurs de direction, Si `false` : cache les FPS, le nombre de particule et les vecteurs de direction |
| bool | GravityEnabled |Si `true` : active la gravité, Si `false` : désactive la gravité|
| bool | OutsideForbidden |Si `true` : active la mort au dela de l'écran, Si `false` : désactive la mort au dela de l'écran|
| bool | LifetimeEnabled |Si `true` : active la durée de vie, Si `false` : désactive la durée de vie|
| bool | DecayEnabled |Si `true` : active le changement des propriétés des particules au cours du temps, Si `false` : désactive le changement des propriétés des particules au cours du temps|
| bool | RangeValue |Si `true` : active l'intervale de valeur pour les propriétés (exemple : couleur entre rouge et bleu), Si `false` : désactive l'intervale de valeur pour les propriétés |
#### Sélection numéro 2 (Les int)
| Type | Nom | Description | Extension
| ---- |---- | ---- | --- |
| int | SpawnX | Définie la position X des particules (si le SpawnType est en "default") | Spawn |
| int | SpawnY | Définie la position Y des particules (si le SpawnType est en "default") | Spawn |
#### Sélection numéro 3 (Les float64)
| Type | Nom | Description | Extension
| ---- |---- | ---- | --- |
| float64 | ShapeSize | Définie la forme dans laquel les particules apparaissent | Forme |
| float64 | SpawnRate | Définie le nombre particule à faire apparaitre par frame | Spawn |
| float64 | GravityForceX | Définie la force de gravité en X | Gravité |
| float64 | GravityForceY | Définie la force de gravité en Y | Gravité |
| float64 | LifetimeMax | Définie la durée de vie maximum d'une particule (en nombre de frame) | Durée de vie |
| float64 | LifetimeMin | Définie la durée de vie minimum d'une particule (en nombre de frame) | Durée de vie |
| float64 | DecayTime | Définie le nombre de frame pour atteindre les objectifs définis | Dégradation |
| float64 | DecayOpacity | Définie l'objectif d'opacité | Dégradation |
| float64 | DecayScaleX | Définie l'objectif de longueur (X) | Dégradation |
| float64 | DecayScaleY | Définie l'objectif de hauteur (Y) | Dégradation |
| float64 | DecayColorR | Définie l'objectif de couleur rouge | Dégradation |
| float64 | DecayColorG | Définie l'objectif de couleur vert | Dégradation |
| float64 | DecayColorB | Définie l'objectif de couleur bleu | Dégradation |
| float64 | DecayRotation | Définie l'objectif de rotation | Dégradation |
| float64 | DefaultRotation | Définie la rotation par défaut | Propriétés par défaut |
| float64 | DefaultColorR | Définie la couleur rouge par défaut | Propriétés par défaut |
| float64 | DefaultColorG | Définie la couleur vert par défaut | Propriétés par défaut |
| float64 | DefaultColorB | Définie la couleur bleu par défaut | Propriétés par défaut |
| float64 | DefaultScaleX | Définie longueur (X) par défaut | Propriétés par défaut |
| float64 | DefaultScaleY | Définie hauteur (Y) par défaut | Propriétés par défaut |
| float64 | DefaultOpacity | Définie l'opacité par défaut | Propriétés par défaut |
| float64 | RangeRotation | Augmente ou réduit | Intervale de Propriétés |
| float64 | RangeDirectionX | Augmente ou réduit | Intervale de Propriétés |
| float64 | RangeDirectionY | Augmente ou réduit | Intervale de Propriétés |
| float64 | RangeScaleY | défini l'écart de génération de la hauteur (Y) | Intervale de Propriétés |
| float64 | RangeScaleX | défini l'écart de génération de la longueur (X) | Intervale de Propriétés |
| float64 | RangeColorR | défini l'écart de génération de la couleur rouge | Intervale de Propriétés |
| float64 | RangeColorG | défini l'écart de génération de la couleur vert | Intervale de Propriétés |
| float64 | RangeColorB | défini l'écart de génération de la couleur bleu | Intervale de Propriétés |
| float64 | RangeOpacity | défini l'écart de génération de l'opacité | Intervale de Propriétés |

**ℹ Note**
> - La variable de SpawnRate ne fonctionne qu'avec une valeur positive, si vous mettez la valeur en négatif elle n'en fera apparaitre aucune : le système considère cela comme un 0.
> - Les variables de couleurs et d'opacité fonctionne avec une valeur de 0 à 1, cependant si vous mettez la couleur au delà elle ne considére pas l'excès : si l'opacité est mise à `4.2`, le système considère cela comme `0.2`.

**ℹ Note**
> L'extension de dégradation est un liste d'objectif défini, si l'ont met l'objectif comme quoi la particule doit avoir bleu à 1, l'extension modifiera la couleur bleu actuel de la particule afin d'atteindre 1, et cela fonctionne avec chacun des objectif définis.
#### Sélection numéro 4 (Méthode de spawn)
| Type | Nom | Description |
| ---- | ---- | ---- |
| string | default | Les particules apparaitront au coordonnées saisi dans SpawnX et SpawnY |
| string | middle | Les particules apparaitront au milieu de l'écran |
| string | random | Les particules apparaitront de manière aléatoire |
| string | cursor | Les particules apparaitront sur le curseur de la souris |
#### Sélection numéro 5 (Forme)
| Type | Nom | Description |
| ---- | ---- | ---- |
| string | default | Les particules apparaitront sans une forme définie |
| string | circle | Les particules apparaitront en forme de cercle (pour un meilleur résultat, mettre le SpawnRate haut est conseillé) |
| string | triangle | Les particules apparaitront en forme de triangle (pour un meilleur résultat, mettre le SpawnRate haut est conseillé) |
| string | square | Les particules apparaitront en forme de carré (pour un meilleur résultat, mettre le SpawnRate haut est conseillé) |
## Liste des variables non dynamique
| Type | Nom | Description |
| ---- | ---- | ---- |
| int | InitNumParticles | définie le nombre de particule qui apparaissent à l'exécution de la fenêtre |
| int | WindowSizeY | définie la hauteur de la fenêtre |
| int | WindowSizeX | définie la longueur de la fenêtre |
| string | WindowsTitle | définie le nom de la fenêtre |
| string | ParticleImage | définie l'image qui sera utilisé comme particule |
## Faire les tests
Pour effectuer des tests, il suffit de se placer (depuis la racine) dans `ParticleSystem/src/particles` et ensuite d'exécuté la commande `go test`

**⚠ Attention**
> Il ne faut pas exécuter ni les tests, ni la fenêtre sur Windows, sinon certaines choses risque de ne pas fonctionner correctement
