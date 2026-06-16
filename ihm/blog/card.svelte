<svelte:options customElement="blog-card" />

<script>
    export let article;
</script>

<div class="card blog-card">
    <div class="card-body">
        <div class="flex">
            <h3 class="card-title">{article.titre}</h3>
            <span class="card-date">{article.date}</span>
        </div>
        <p class="card-text">{article.description}</p>
        <div class="read-more">
            <span class="link">Lire la suite <i class="bi bi-arrow-right"></i></span>
        </div>
    </div>
</div>

<style>
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");
    @import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css");

    .blog-card {
        width: 100%;
        max-width: 800px;
        margin-bottom: 2rem;
        transition: transform 0.2s, box-shadow 0.2s;
        cursor: pointer;
        border: none;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        border-radius: 8px;
    }

    .blog-card:hover {
        transform: translateY(-5px);
        box-shadow: 0 8px 15px rgba(0,0,0,0.15);
    }

    .card-title {
        color: #333;
        font-weight: bold;
        margin-bottom: 1rem;
    }

    .card-text {
        color: #666;
        display: -webkit-box;
        -webkit-line-clamp: 3;
        line-clamp: 3;
        -webkit-box-orient: vertical;
        overflow: hidden;
    }

    .card-date {
        font-size: 0.9rem;
        color: #999;
        margin-bottom: 1rem;
    }

    .read-more {
        margin-top: 1rem;
        text-align: right;
    }

    .link {
        color: var(--primary-color, #81c553);
        font-weight: bold;
        text-decoration: none;
        display: inline-flex;
        align-items: center;
        gap: 0.5rem;
    }

    .link i {
        transition: transform 0.3s ease;
    }

    .blog-card:hover .link i {
        transform: translateX(5px);
    }
</style>