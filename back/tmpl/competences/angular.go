package tmpl

// AngularCompetence contient la description HTML de la compétence Angular.
const AngularCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        Angular est un framework open-source basé sur TypeScript, maintenu par l'équipe Angular de Google. 
        Il est conçu pour créer des applications web monopages (Single Page Applications) complexes et performantes. 
        Angular repose sur une architecture par composants et propose un ensemble complet d'outils pour le routage, la gestion d'état et la communication avec les API.
    </p>
    <p>
        Mon expérience avec Angular inclut la création de composants modulaires, la gestion des services et de l'injection de dépendances, ainsi que 
        l'utilisation du routage pour proposer une navigation fluide. C'est un framework robuste, particulièrement adapté aux applications d'entreprise 
        où la maintenabilité et la scalabilité sont essentielles.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        J'ai appliqué mes compétences Angular sur le développement de Project Management Tool (PMT), une application web full-stack où Angular était au cœur du front-end.
    </p>
    
    <p>
        <strong>SPA modulable et maintenable (<a class="link" href="/projets/detail.html?id=5">PMT</a>).</strong>
        <br>Pour le projet PMT, j'ai structuré le front-end comme une Single-Page Application dynamique, avec une architecture par composants (Dashboard, Projet, Tâche). 
        Cela a permis de créer une interface cohérente, facile à maintenir et évolutive, tout en facilitant la gestion de projets complexes.
    </p>

    <p>
        <strong>Centralisation de l'état et communication avec le backend (<a class="link" href="/projets/detail.html?id=5">PMT</a>).</strong>
        <br>Les services Angular ont été utilisés pour gérer la logique métier et la communication asynchrone avec l'API backend 
        (<a class="link" href=/competences/detail.html?id=7>Spring</a>). 
        Cette organisation a permis de découpler les composants de la logique de récupération des données, garantissant un code clair et une application fiable, 
        avec un flux de données optimisé.
    </p>

    <p>
        <strong>Intégration full-stack et environnement conteneurisé (<a class="link" href="/projets/detail.html?id=5">PMT</a>).</strong>
        <br>Le front-end Angular a été intégré à une API RESTful et déployé dans un environnement Docker, couvrant l'ensemble du cycle d'intégration, 
        de la définition des contrats de données jusqu'à la gestion des requêtes HTTP. Cette approche a permis d'assurer une cohésion parfaite entre les 
        différentes couches de l'application et un déploiement fluide.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Mon niveau de pratique d'Angular me permet de concevoir des applications front-end basiques, fonctionnelles et maintenables, mais je ne me considère pas encore comme un expert. 
        Angular est un framework puissant et complet, mais sa complexité implique pas mal de pratique avant de le maîtriser completement. 
        Sur le projet PMT, j'ai beaucoup appris sur la structuration d'une SPA, la gestion de services et l'intégration avec un backend, mais il me reste encore 
        à découvrir avant de pouvoir tirer pleinement parti de toutes ses possibilités.
    </p>
    <p>
        Pour l'instant, je préfère me concentrer sur des frameworks plus légers et réactifs comme Svelte, qui me permettent de créer rapidement des interfaces performantes tout 
        en restant proche des fondamentaux du web. Angular reste un outil que je peux utiliser et approfondir si le projet le nécessite.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À court terme, je n'ai pas de plan pour devenir un expert Angular, mais je reste curieux des nouvelles fonctionnalités comme les Standalone Components ou le Server-Side Rendering (SSR). 
        J'aimerais continuer à en apprendre davantage progressivement, au fil des projets où Angular s'avère pertinent.
    </p>
    <p>
        Mon intérêt actuel se porte davantage sur des frameworks légers et modernes qui facilitent le développement rapide et la réactivité des interfaces. 
        Cependant, je garde Angular dans ma boîte à outils pour les projets d'envergure qui nécessitent sa robustesse et sa structure complète.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=5">PMT</a></li>
        </ul>
    </div>
</section>
`
