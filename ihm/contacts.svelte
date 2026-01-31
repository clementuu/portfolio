<svelte:options customElement="contacts-portfolio" />

<svelte:head>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css">
</svelte:head>

<script>
    import Header from "./elements/header.svelte";

    const contacts = [
        {
            type: "Téléphone",
            value: "+33 6 25 45 25 66",
            link: "tel:+33625452566",
            iconClass: "bi-telephone"
        },
        {
            type: "Email",
            value: "clementcalia@gmail.com",
            link: "mailto:clementcalia@gmail.com",
            iconClass: "bi-envelope"
        },
        {
            type: "LinkedIn",
            value: "Clément Calia",
            link: "https://www.linkedin.com/in/clement-calia",
            iconClass: "bi-linkedin"
        },
        {
            type: "GitHub",
            value: "clementuu",
            link: "https://github.com/clementuu",
            iconClass: "bi-github"
        }
    ];
</script>

<Header/>

<div class="contact-div">
    <h1><b>Me Contacter</b></h1>
    <div class="contact-list">
        {#each contacts as contact}
            <div class="contact-item">
                <div class="contact-icon">
                    <i class="bi {contact.iconClass}"></i>
                </div>
                <strong>{contact.type}:</strong>
                <a class="link" href={contact.link} target="_blank" rel="noopener noreferrer">{contact.value}</a>
            </div>
        {/each}
    </div>
</div>

<style>
    @import "./style/style.css";
	@import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css");
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .contact-list {
        display: flex;
        flex-direction: column;
        gap: 1em;
        padding: 1em;
        max-width: 600px;
        margin: auto;
    }

    .contact-item {
        display: flex;
        align-items: center;
        gap: 1em;
        padding: 0.5em 0;
        border-bottom: 1px solid #eee;
    }

    .contact-item:last-child {
        border-bottom: none;
    }

    .contact-icon {
        width: 24px;
        height: 24px;
    }

    .contact-item strong {
        flex-grow: 1;
    }

    .contact-div {
        position: relative;
        top: 0;
    }

    @media (min-width: 768px) {
		.contact-div {
			top: var(--header-height);
		}
	}
</style>