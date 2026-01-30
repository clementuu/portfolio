package tmpl

// SQLCompetence contient la description HTML de la compétence SQL.
const SQLCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        SQL (Structured Query Language) est le langage de référence pour interagir avec les bases de données 
        relationnelles. Dans un contexte professionnel où les organisations s'appuient de plus en plus sur 
        la donnée pour piloter leurs décisions, optimiser leurs processus et automatiser leurs systèmes, 
        la maîtrise de SQL est devenue indispensable. L'actualité technologique - montée en puissance de la 
        data, importance du reporting en temps réel, exigences de conformité et de qualité des données - 
        renforce encore la nécessité de disposer de bases solides en modélisation et manipulation de données.
    </p>
    <p>
        Mon expérience avec SQL couvre l'ensemble du cycle de vie des données : conception de schémas, 
        optimisation de requêtes, migrations critiques, indexation, gestion de transactions et intégration 
        dans des architectures applicatives. J'ai travaillé avec plusieurs SGBD, notamment 
        et SQLite, en utilisant SQL aussi bien directement qu'à travers des ORM ou des 
        bibliothèques bas niveau. Cette compétence est au cœur de mes projets, qu'il s'agisse d'applications 
        web, de systèmes embarqués ou de plateformes métier complexes.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        Ma maîtrise de SQL s'est construite à travers des projets concrets où j'ai dû concevoir, refondre, 
        optimiser et sécuriser des bases de données complètes. Voici trois anecdotes illustrant l'application 
        directe de cette compétence.
    </p>

    <p>
        <strong>Refonte de schémas complexes et modernisation des données (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>, <a class="link" href="/projets/detail.html?id=6">PMT</a>).</strong>
        <br>Sur les projets Escarcelle et PMT, j'ai conçu et fait évoluer des 
        schémas relationnels complets. J'ai notamment transformé un système de catégories basé sur l'ordre 
        des IDs en une structure hiérarchique moderne utilisant un <i>parent_id</i>, et j'ai étendu les tables 
        produits pour intégrer un système multi-code-barres flexible. Résultat une base plus cohérente, plus performante 
        et capable de supporter de nouvelles fonctionnalités métier. J'ai apporté une vision structurée et évolutive de la donnée, 
        permettant de moderniser un socle historique sans rupture.
    </p>

    <p>
        <strong>Migration critique de MySQL 5 vers MySQL 8 (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>).</strong>
        <br>J'ai mené une migration complète du SGBD, incluant l'analyse d'impact, la correction de requêtes 
        dépréciées, la mise à jour des index et l'adaptation du code applicatif. Cette opération était 
        essentielle pour garantir la sécurité, la performance et la pérennité du système. La transition s'est déroulée de manière totalement fluide, 
        sans perte de données ni interruption de service. J'ai apporté une réelle valeur en sécurisant une opération à fort enjeu technique et métier, 
        tout en améliorant la performance globale du système.
    </p>

    <p>
        <strong>Gestion d'une base locale embarquée et synchronisation (<a class="link" href="/projets/detail.html?id=2">Caisse</a>).</strong>
        <br>Pour l'application de Caisse, j'ai conçu et géré une base SQLite 
        locale, incluant la gestion des migrations et un mécanisme de synchronisation 
        bidirectionnelle avec l'API centrale. Cette architecture a permis à l'application de fonctionner hors ligne tout en garantissant 
        la cohérence des données lors de la resynchronisation. Ma contribution a consisté à fournir une solution robuste à un problème 
        critique de continuité de service, essentielle pour un outil utilisé quotidiennement en production.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Je considère ma maîtrise de SQL comme avancée et opérationnelle. C'est une compétence 
        essentielle dans mon métier de développeur full-stack, car elle conditionne la qualité, la performance 
        et la fiabilité des applications que je conçois. SQL occupe une place centrale dans mon profil 
        d'ingénierie : elle me permet de comprendre en profondeur les enjeux liés à la donnée, d'optimiser 
        les performances et de garantir la cohérence des systèmes.
    </p>
    <p>
        J'ai acquis cette compétence rapidement grâce à une pratique intensive sur des projets réels, 
        notamment des refontes de schémas et des migrations sensibles. Mon conseil : avant de chercher à 
        optimiser des requêtes complexes, il est essentiel de maîtriser les bases - normalisation, 
        indexation, relations, transactions. Une bonne modélisation évite 80 % des problèmes de performance.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À moyen terme, je souhaite approfondir mes compétences en optimisation avancée (indexation fine, 
        analyse de plans d'exécution, partitionnement) et explorer davantage les problématiques de 
        scalabilité des bases relationnelles. Je m'intéresse également aux solutions hybrides combinant 
        SQL et NoSQL pour répondre à des besoins métier variés.
    </p>
    <p>
        Je suis actuellement des ressources sur l'optimisation MySQL et j'envisage de suivre une formation 
        dédiée aux architectures data modernes. Cette montée en compétence s'inscrit dans mon projet 
        professionnel : devenir un expert capable de concevoir des systèmes de données performants, 
        fiables et adaptés aux enjeux actuels.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
