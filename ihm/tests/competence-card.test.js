import { render, screen } from '@testing-library/svelte';
import { describe, it, expect, afterEach, vi } from 'vitest';
import Card from '../competences/card.svelte';

afterEach(() => {
  document.body.innerHTML = '';
  vi.restoreAllMocks();
});

describe('CompetenceCard component', () => {
  const sampleCompetence = {
    id: 1,
    name: 'GoLang',
    image: '/assets/icons8-golang.svg',
    rating: 4,
    type: 'Dev'
  };

  it('renders name, image, rating, and type', () => {
    // Mock window.location.href to prevent actual navigation during test
    const assignMock = vi.fn();
    Object.defineProperty(window, 'location', {
        writable: true,
        value: { assign: assignMock },
    });

    const { container } = render(Card, { props: { competence: sampleCompetence } });

    expect(screen.getByText('GoLang')).toBeInTheDocument();

    const img = container.querySelector('img');
    expect(img).toBeTruthy();
    expect(img).toHaveAttribute('src', '/assets/icons8-golang.svg');
    expect(img).toHaveAttribute('alt', 'GoLang');

    expect(screen.getByText('⭐'.repeat(4) + '☆'.repeat(1))).toBeInTheDocument(); // Check stars

    expect(screen.getByText('#Dev')).toBeInTheDocument();
    expect(screen.getByText('#Dev')).toHaveClass('card-span');
    expect(screen.getByText('#Dev')).toHaveClass('technique');
  });

  const sampleCompetenceDevOps = {
    id: 2,
    name: 'Docker',
    image: '/assets/icons8-docker.svg',
    rating: 3,
    type: 'DevOps'
  };

  it('applies devops class for DevOps type', () => {
    const assignMock = vi.fn();
    Object.defineProperty(window, 'location', {
        writable: true,
        value: { assign: assignMock },
    });

    render(Card, { props: { competence: sampleCompetenceDevOps } });
    expect(screen.getByText('#DevOps')).toHaveClass('devops');
  });

  const sampleCompetenceHumain = {
    id: 3,
    name: 'Créativité',
    image: '/assets/icons8-creativity-80.png',
    rating: 5,
    type: 'Humain'
  };

  it('applies humain class for Humain type', () => {
    const assignMock = vi.fn();
    Object.defineProperty(window, 'location', {
        writable: true,
        value: { assign: assignMock },
    });

    render(Card, { props: { competence: sampleCompetenceHumain } });
    expect(screen.getByText('#Humain')).toHaveClass('humain');
  });
});
