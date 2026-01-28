package tmpl

// SQLComp contient la description HTML de la compétence SQL.
const SQLComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        SQL (Structured Query Language) est le langage standard pour la gestion et la manipulation des 
        bases de données relationnelles. Sa maîtrise est fondamentale pour interagir avec les données, 
        que ce soit pour les lire, les insérer, les mettre à jour ou les supprimer.
    </p>
    <p>
        Mon expérience avec SQL s'étend de la conception de schémas de bases de données (tables, relations, 
        index) à l'écriture de requêtes complexes pour l'analyse de données et le reporting. J'ai travaillé 
        avec plusieurs systèmes de gestion de bases de données, notamment MySQL et SQLite, en utilisant SQL 
        aussi bien directement que via des ORM ou des bibliothèques bas niveau.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        SQL est une compétence centrale que j'ai appliquée dans la majorité de mes projets pour la modélisation, la manipulation et l'optimisation des données. Mon expérience couvre à la fois des systèmes de gestion de bases de données client-serveur comme <strong>MySQL</strong> et des bases de données embarquées comme <strong>SQLite</strong>.
    </p>
    <ul>
        <li>
            <strong>Modélisation et refonte de schémas :</strong> Dans les projets <strong>PMT</strong> et <strong>Escarcelle</strong>, j'ai conçu des schémas de bases de données relationnelles et mené des opérations de refonte complexes. J'ai notamment fait évoluer un système de catégories de produits vers une structure hiérarchique plus moderne et flexible, et étendu des tables existantes pour supporter de nouvelles fonctionnalités métier comme le multi-code-barres.
        </li>
        <li>
            <strong>Migration de base de données :</strong> J'ai mené à bien une migration complète de <strong>MySQL 5 à MySQL 8</strong> pour le projet Escarcelle. Cette mission critique a inclus une analyse d'impact, la réécriture de requêtes dépréciées et l'adaptation du code applicatif pour assurer une transition sans perte de données.
        </li>
        <li>
            <strong>Gestion de base de données locale et synchronisation :</strong> Pour l'application de <strong>Caisse</strong>, j'ai mis en place et géré une base de données <strong>SQLite</strong> locale, incluant la gestion des migrations, l'indexation pour la performance, et un mécanisme de synchronisation bidirectionnelle avec l'API centrale.
        </li>
    </ul>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=2">Caisse</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT</a></li>
        </ul>
    </div>
</section>
`
