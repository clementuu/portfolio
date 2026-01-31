package tmpl

// BashCompetence contient la description HTML de la compétence Bash.
const BashCompetence = `
<section class="comp-section">
    <h3><b>Définition</b></h3>
    <p>
        Bash (Bourne-Again SHell) est l'interpréteur de commandes par défaut sur la plupart des systèmes Linux. Au-delà d'une simple interface en ligne de commande, 
        c'est un langage de script puissant qui permet d'automatiser une grande variété de tâches système et de développement.
    </p>
    <p>
        Utilisant Linux comme environnement de travail principal depuis plus de deux ans, j'ai rapidement adopté Bash comme un outil essentiel de mon quotidien. 
        Je l'utilise aussi bien pour des commandes simples que pour l'écriture de scripts d'automatisation. Sa maîtrise est cruciale pour l'efficacité en DevOps et 
        l'administration système.
    </p>
</section>

<section class="comp-section">
    <h3><b>Éléments de preuve</b></h3>
    <p>
        Mon utilisation de Bash va de la simple gestion de mon environnement de développement à la création de scripts pour des opérations complexes et critiques.
    </p>

    <p>
        <strong>Automatisation de migration de base de données (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>).</strong>
        <br>J'ai développé un script Bash pour automatiser entièrement la migration d'une base de données de <strong>MySQL 5 vers MySQL 8</strong>
        pour le projet Escarcelle. Ce script intégrait des points de contrôle à chaque étape cruciale (sauvegarde, mise à jour, redémarrage du service) 
        et générait un fichier de log détaillé pour assurer la traçabilité et la sécurité de l'opération. Le résultat a été une migration de base de données complexe 
        réalisée de manière fiable et sans erreur, réduisant considérablement le temps d'arrêt. Ma valeur ajoutée a été de transformer 
        une tâche manuelle risquée en un processus automatisé et sécurisé.
    </p>

    <p>
        <strong>Gestion quotidienne et optimisation de l'environnement de dev (<a class="link" href="/projets/detail.html?id=1">Escarcelle</a>, <a class="link" href="/projets/detail.html?id=6">PMT</a>, <a class="link" href="/projets/detail.html?id=2">Caisse</a>).</strong>
        <br>J'utilise couramment Bash pour naviguer dans le système de fichiers, gérer mes projets avec Git, manipuler des conteneurs Docker et lancer mes applications. 
        J'ai également créé des alias et des fonctions personnalisées dans mon <code>.bashrc</code> pour accélérer les tâches répétitives. 
        Le résultat est un environnement de développement hautement productif et personnalisé. Ma valeur ajoutée a été d'optimiser mon flux de travail quotidien, 
        me permettant de me concentrer davantage sur le développement.
    </p>
</section>

<section class="comp-section">
    <h3><b>Autocritique</b></h3>
    <p>
        Ma maîtrise de Bash est avancée. Je suis capable d'écrire des scripts complexes et robustes pour automatiser diverses tâches. 
        Je comprends bien les concepts de redirection, de piping, de gestion des variables et des fonctions en Bash. C'est une compétence pratique que 
        j'ai acquise par l'usage quotidien et la résolution de problèmes spécifiques.
    </p>
    <p>
        Un conseil pour ceux qui veulent améliorer leur Bash : lisez des scripts existants, pratiquez régulièrement et n'hésitez pas à consulter le manuel 
        ou des ressources en ligne. L'automatisation des tâches répétitives vous fera gagner un temps précieux.
    </p>
</section>

<section class="comp-section">
    <h3><b>Évolution</b></h3>
    <p>
        À moyen terme, je souhaite approfondir ma connaissance des outils de gestion de configuration comme Ansible ou Terraform, 
        qui s'appuient souvent sur des scripts shell pour l'automatisation. Mon objectif est d'intégrer Bash de manière 
        plus systématisée dans des pipelines CI/CD complexes et des infrastructures cloud.
    </p>
    <p>
        Je prévois de suivre des tutoriels sur Ansible et de créer des playbooks pour automatiser le déploiement de mes applications.
    </p>
</section>

<section class="comp-section">
    <h5>Projets Associés</h5>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
        </ul>
    </div>
</section>
`
