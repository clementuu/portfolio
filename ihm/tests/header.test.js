import { render, screen } from '@testing-library/svelte';
import { describe, it, expect, afterEach } from 'vitest';
import Header from '../elements/header.svelte';

afterEach(() => {
  document.body.innerHTML = '';
});

describe('Header - static checks (no network)', () => {
  it('renders main links and structure', () => {
    const { container } = render(Header);
    expect(screen.getByText('Clément Calia')).toBeInTheDocument();
    expect(screen.getByText('À propos')).toBeInTheDocument();
    expect(screen.getByText('Projets')).toBeInTheDocument();
    expect(screen.getByText('Compétences')).toBeInTheDocument();
    expect(screen.getByText('Contacts')).toBeInTheDocument();

    const nav = container.querySelector('nav.header');
    expect(nav).toBeTruthy();
  });

  it('renders competence dropdown static items', () => {
    render(Header);
    expect(screen.getByText('Développement')).toBeInTheDocument();
    expect(screen.getByText('DevOps')).toBeInTheDocument();
    expect(screen.getByText('Soft skills')).toBeInTheDocument();
  });

  it('links have correct hrefs for static items', () => {
    render(Header);
    const accueil = screen.getByText('Clément Calia').closest('a');
    expect(accueil).toHaveAttribute('href', '/');

    const contacts = screen.getByText('Contacts').closest('a');
    expect(contacts).toHaveAttribute('href', '/contacts.html');
  });
});
