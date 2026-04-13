<svelte:options customElement="contacts-portfolio" />

<svelte:head>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css">
</svelte:head>

<script>
    import Header from "./elements/header.svelte";

    let verified = false;
    let showCaptcha = false;
    let captchaA = 0;
    let captchaB = 0;
    let userAnswer = "";

    function generateCaptcha() {
        captchaA = Math.floor(Math.random() * 10);
        captchaB = Math.floor(Math.random() * 10);
        showCaptcha = true;
    }

    function verifyCaptcha() {
        if (parseInt(userAnswer) === captchaA + captchaB) {
            verified = true;
            showCaptcha = false;
        } else {
            alert("Mauvaise réponse. Veuillez réessayer.");
            generateCaptcha();
            userAnswer = "";
        }
    }

    const contacts = [
        {
            type: "Téléphone",
            // "+33 6 25 45 25 66"
            encodedValue: "KzMzIDYgMjUgNDUgMjUgNjY=",
            // "tel:+33625452566"
            encodedLink: "dGVsOiszMzYyNTQ1MjU2Ng==",
            iconClass: "bi-telephone",
            protected: true
        },
        {
            type: "Email",
            // "clementcalia@gmail.com"
            encodedValue: "Y2xlbWVudGNhbGlhQGdtYWlsLmNvbQ==",
            // "mailto:clementcalia@gmail.com"
            encodedLink: "bWFpbHRvOmNsZW1lbnRjYWxpYUBnbWFpbC5jb20=",
            iconClass: "bi-envelope",
            protected: true
        },
        {
            type: "LinkedIn",
            value: "Clément Calia",
            link: "https://www.linkedin.com/in/clement-calia",
            iconClass: "bi-linkedin",
            protected: false
        },
        {
            type: "GitHub",
            value: "clementuu",
            link: "https://github.com/clementuu",
            iconClass: "bi-github",
            protected: false
        }
    ];

    function decode(str) {
        try {
            return atob(str);
        } catch (e) {
            return "";
        }
    }
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
                
                {#if !contact.protected}
                    <a class="link" href={contact.link} target="_blank" rel="noopener noreferrer">{contact.value}</a>
                {:else if verified}
                    <a class="link" href={decode(contact.encodedLink)} target="_blank" rel="noopener noreferrer">
                        {decode(contact.encodedValue)}
                    </a>
                {:else}
                    <button class="btn btn-sm btn-outline-secondary" on:click={generateCaptcha}>
                        Révéler
                    </button>
                {/if}
            </div>
        {/each}
    </div>

    {#if showCaptcha}
        <div class="captcha-overlay">
            <div class="captcha-modal">
                <h5>Vérification</h5>
                <p>Pour lutter contre le spam, veuillez résoudre ce calcul simple :</p>
                <div class="captcha-math">
                    {captchaA} + {captchaB} = 
                    <input type="number" bind:value={userAnswer} on:keydown={(e) => e.key === 'Enter' && verifyCaptcha()} />
                </div>
                <div class="captcha-actions">
                    <button class="btn captcha-btn" on:click={verifyCaptcha}>Valider</button>
                    <button class="btn btn-link text-secondary" on:click={() => showCaptcha = false}>Annuler</button>
                </div>
            </div>
        </div>
    {/if}
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

    .captcha-overlay {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0,0,0,0.5);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
    }

    .captcha-modal {
        background: white;
        padding: 2em;
        border-radius: 8px;
        box-shadow: 0 4px 15px rgba(0,0,0,0.2);
        max-width: 400px;
        text-align: center;
    }

    .captcha-math {
        font-size: 1.5em;
        margin: 1em 0;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 0.5em;
    }

    .captcha-math input {
        width: 60px;
        padding: 0.2em;
        text-align: center;
    }

    .captcha-actions {
        display: flex;
        flex-direction: column;
        gap: 0.5em;
    }

    .captcha-btn {
        background-color: var(--primary-color);
        color: white;
    }

    .captcha-btn:hover {
        filter: brightness(1.1);
    }

    @media (min-width: 800px) {
		.contact-div {
			top: var(--header-height);
		}
	}
</style>