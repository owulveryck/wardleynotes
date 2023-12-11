## Table de doctrine
La Table de Doctrine de Simon Wardley est un cadre utilisé dans la stratégie d'entreprise et le développement organisationnel, qui se concentre sur la manière dont les organisations peuvent naviguer et s'adapter dans un paysage en constante évolution. 

Pour les _Wardley Mappers_, la Doctrine n'est pas un terme effrayant mais plutôt un ensemble de principes universellement applicables. 
Des principes qui fonctionnent dans toutes les situations, pour chaque organisation.

Pensez-y. Est-ce qu'une personne raisonnable pourrait soutenir que des principes tels que "_Connaître vos utilisateurs_" ou "_Remettre en question les hypothèses_" devraient être ignorés ?

Cela n'aurait tout simplement pas de sens. Cela les rend universellement applicables.

Chaque organisation devrait intégrer ces principes.

Malheureusement, elles ne le font souvent pas, car elles ne sont pas conscientes que c'est quelque chose qu'elles devraient gérer.

Simon Wardley a identifié quarante principes universels.

Un constat est que les organisations faibles en doctrine sont également faibles en affaires, car elles ne peuvent pas réagir aux changements du marché, même si leur activité existante paraît bonne.

### Découpage

La doctrine est divisée en quatre étapes disctinctes:

Chaque étape s'appuie sur la précédente, formant une approche intégrée du développement organisationnel et de la planification stratégique. 
Cette méthode est particulièrement pertinente dans l'environnement commercial actuel, rapide et axé sur la technologie, où l'adaptabilité et l'agilité sont essentielles au succès.

### 1. Élimination des Pratiques Contreproductives

#### Focus
Cette étape vise à éliminer les comportements et habitudes qui limitent la capacité de l'organisation à s'adapter et à évoluer.
#### Actions Principales
Suppression des silos organisationnels et des obstacles internes.
Abandon des pratiques inefficaces ou contre-productives.
Promotion de la communication transparente et de la collaboration.

#### Objectif 
Créer une fondation robuste en traitant et en corrigeant les problèmes internes qui freinent le progrès.

### 2. Alignement avec le Contexte de l'Environnement

#### Focus
Cette étape met l'accent sur la compréhension du contexte dans lequel l'organisation opère.

#### Actions Principales
Réalisation d'études de marché détaillées.
Compréhension des besoins et attentes des clients.
Reconnaissance du paysage concurrentiel et des tendances technologiques.

#### Objectif 
Acquérir une compréhension précise des facteurs externes influençant l'organisation pour mieux adapter les stratégies aux défis rencontrés.

### 3. Optimisation des Ressources pour une Efficacité Accrue

#### Focus 
Cette étape se concentre sur l'amélioration de l'efficacité et de la productivité.

#### Actions Principales
Rationalisation des processus.
Adoption de principes Lean.
Utilisation de la technologie et de l'innovation pour réaliser plus avec moins de ressources.

#### Objectif
Optimiser les opérations pour offrir une plus grande valeur à un coût réduit, améliorant ainsi la performance globale.

### 4. Adaptabilité et Évolution Continues

#### Focus
La dernière étape consiste à encourager une culture d'amélioration et d'adaptation continue.

#### Actions Principales
Stimuler l'innovation et l'expérimentation.
Révision et mise à jour régulières des stratégies.
Adaptation rapide aux changements du marché et technologiques.

#### Objectif 
Assurer la durabilité et le succès à long terme en restant agile et flexible, prêt à évoluer selon les besoins.

## Self-assessment

Inspiré de l'outil de [Krzysztof Daniel.](https://doctrine.wardleymaps.com/), [cette grille](https://owulveryck.github.io/wardleynotes/Refs/doctrine_fr.html) permet d'évaluer la position ressentie d'une entreprises sur chaque phase et par pillier:
- communication
- développement
- opération
- apprentissage
- direction
- structure

Cette grille utilise propose d'évaluer la **perception d'une équipe** sur un principe spécifique en utilisant un système simple de trois couleurs pour représenter les sentiments ou les opinions :

* Vert : "Nous le faisons très bien." Cette couleur indique que l'équipe se sent confiante et performante par rapport à ce principe. Ils considèrent qu'ils appliquent le principe de manière efficace et réussie.

* Orange : "Nous pourrions améliorer certaines choses." Cette couleur suggère que, bien que l'équipe mette en œuvre le principe, il y a des domaines d'amélioration. Cela peut signifier qu'ils reconnaissent les défis ou les lacunes dans leur approche actuelle.

* Rouge : "Nous ne le faisons pas." Cette couleur indique que le principe n'est pas actuellement mis en pratique ou est considéré comme un domaine négligé. Cela peut signaler un besoin d'attention immédiate ou de changements majeurs.

Ce système de couleurs permet une évaluation rapide et intuitive de l'attitude de l'équipe envers un principe donné, sans nécessiter un modèle de maturité complexe ou une échelle détaillée.

### Aggrégation

Il est possible d'exporter le résultat au format `json`.

L'ensemble des grilles peuvent être agrégées grâce à l'outil suivant en Go: 
`cat *.json | go run main.go > resultat.json`

le fichier `resultat.json` peut ensuite être restauré dans la même page pour avoir une vue consolidée du résultat.

