import { render, screen, fireEvent } from '@testing-library/svelte';
import { describe, it, expect, afterEach } from 'vitest';
import Card from '../projets/card.svelte'; // adapte le chemin si besoin

afterEach(() => {
  document.body.innerHTML = '';
});

describe('Card component - rendu et interactions simples', () => {
  const sampleProjet = {
    id: 123,
    name: 'Projet Gamma',
    image: '/images/gamma.png',
    resume: 'Un super projet',
    competences: [
      { name: 'JavaScript', type: 'Dev' },
      { name: 'CI/CD', type: 'DevOps' },
      { name: 'Communication', type: 'Humain' }
    ]
  };

  it('rend le nom, le résumé et l\'image avec les bons attributs', () => {
    const { container } = render(Card, { props: { projet: sampleProjet } });

    // titre
    expect(screen.getByText('Projet Gamma')).toBeInTheDocument();

    // résumé
    expect(screen.getByText('Un super projet')).toBeInTheDocument();

    // image : src et alt
    const img = container.querySelector('img');
    expect(img).toBeTruthy();
    expect(img).toHaveAttribute('src', '/images/gamma.png');
    expect(img).toHaveAttribute('alt', 'Projet Gamma');
  });

  it('rend les badges de compétences avec les classes attendues', () => {
    render(Card, { props: { projet: sampleProjet } });

    const badges = Array.from(document.querySelectorAll('.competences .card-span'));
    expect(badges.length).toBe(3);

    // Vérifie le mapping des classes
    expect(badges[0].textContent).toBe('JavaScript');
    expect(badges[0].classList.contains('technique')).toBe(true);

    expect(badges[1].textContent).toBe('CI/CD');
    expect(badges[1].classList.contains('devops')).toBe(true);

    expect(badges[2].textContent).toBe('Communication');
    expect(badges[2].classList.contains('humain')).toBe(true);
  });
});
