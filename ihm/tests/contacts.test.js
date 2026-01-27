import { render, screen } from '@testing-library/svelte';
import { describe, it, expect, afterEach } from 'vitest';
import Contacts from '../contacts.svelte';

afterEach(() => {
  document.body.innerHTML = '';
});

describe('Contacts component - static checks', () => {
  it('renders the heading Me Contacter', () => {
    render(Contacts);
    expect(screen.getByText('Me Contacter')).toBeInTheDocument();
  });

  it('renders four contact items', () => {
    const { container } = render(Contacts);
    const items = container.querySelectorAll('.contact-item');
    expect(items.length).toBe(4);
  });

  it('renders each contact type label and value', () => {
    render(Contacts);

    // Types
    expect(screen.getByText('Téléphone:')).toBeInTheDocument();
    expect(screen.getByText('Email:')).toBeInTheDocument();
    expect(screen.getByText('LinkedIn:')).toBeInTheDocument();
    expect(screen.getByText('GitHub:')).toBeInTheDocument();

    // Values / links
    expect(screen.getByText('+33 6 25 45 25 66')).toBeInTheDocument();
    expect(screen.getByText('clementcalia@gmail.com')).toBeInTheDocument();
    expect(screen.getByText('Clément Calia')).toBeInTheDocument();
    expect(screen.getByText('clementuu')).toBeInTheDocument();
  });

  it('anchors have correct href, target and rel attributes', () => {
    render(Contacts);

    const telLink = screen.getByText('+33 6 25 45 25 66').closest('a');
    expect(telLink).toHaveAttribute('href', 'tel:+33625452566');
    expect(telLink).toHaveAttribute('target', '_blank');
    expect(telLink).toHaveAttribute('rel', 'noopener noreferrer');

    const mailLink = screen.getByText('clementcalia@gmail.com').closest('a');
    expect(mailLink).toHaveAttribute('href', 'mailto:clementcalia@gmail.com');

    const linkedinLink = screen.getByText('Clément Calia').closest('a');
    expect(linkedinLink).toHaveAttribute('href', 'https://www.linkedin.com/in/clement-calia');

    const githubLink = screen.getByText('clementuu').closest('a');
    expect(githubLink).toHaveAttribute('href', 'https://github.com/clementuu');
  });

  it('renders icon elements with expected classes', () => {
    const { container } = render(Contacts);
    const icons = container.querySelectorAll('.contact-icon i');

    // Vérifie que chaque icône contient la classe bi et une classe spécifique
    expect(icons.length).toBe(4);
    expect(icons[0].classList.contains('bi')).toBe(true);
    expect(icons[0].classList.contains('bi-telephone')).toBe(true);

    expect(icons[1].classList.contains('bi-envelope')).toBe(true);
    expect(icons[2].classList.contains('bi-linkedin')).toBe(true);
    expect(icons[3].classList.contains('bi-github')).toBe(true);
  });
});
