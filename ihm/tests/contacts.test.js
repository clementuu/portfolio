import { render, screen, fireEvent } from '@testing-library/svelte';
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

  it('renders each contact type label and handles protected fields', () => {
    render(Contacts);

    // Types labels should always be present
    expect(screen.getByText('Téléphone:')).toBeInTheDocument();
    expect(screen.getByText('Email:')).toBeInTheDocument();
    expect(screen.getByText('LinkedIn:')).toBeInTheDocument();
    expect(screen.getByText('GitHub:')).toBeInTheDocument();

    // Protected fields should show "Révéler" initially
    const revealButtons = screen.getAllByText('Révéler');
    expect(revealButtons).toHaveLength(2);

    // Public fields should be visible
    const linkedinLinks = screen.getAllByText('Clément Calia');
    expect(linkedinLinks).toHaveLength(2);
    expect(screen.getByText('clementuu')).toBeInTheDocument();
    
    // Sensitive values should NOT be in the document initially
    expect(screen.queryByText('+33 6 25 45 25 66')).not.toBeInTheDocument();
    expect(screen.queryByText('clementcalia@gmail.com')).not.toBeInTheDocument();
  });

  it('public anchors have correct href, target and rel attributes', () => {
    render(Contacts);

    const linkedinLinks = screen.getAllByText('Clément Calia');
    const linkedinLink = linkedinLinks[1].closest('a');
    expect(linkedinLink).toHaveAttribute('href', 'https://www.linkedin.com/in/clement-calia');

    const githubLink = screen.getByText('clementuu').closest('a');
    expect(githubLink).toHaveAttribute('href', 'https://github.com/clementuu');
  });

  it('renders icon elements with expected classes', () => {
    const { container } = render(Contacts);
    const icons = container.querySelectorAll('.contact-icon i');

    expect(icons.length).toBe(4);
    expect(icons[0].classList.contains('bi-telephone')).toBe(true);
    expect(icons[1].classList.contains('bi-envelope')).toBe(true);
    expect(icons[2].classList.contains('bi-linkedin')).toBe(true);
    expect(icons[3].classList.contains('bi-github')).toBe(true);
  });
});

describe('Contacts component - captcha interaction', () => {
  it('shows captcha when clicking Reveal', async () => {
    render(Contacts);
    const revealButton = screen.getAllByText('Révéler')[0];
    await fireEvent.click(revealButton);
    
    expect(screen.getByText(/Pour lutter contre le spam/)).toBeInTheDocument();
    expect(screen.getByText('Vérification')).toBeInTheDocument();
  });

  it('cancels captcha when clicking Annuler', async () => {
    render(Contacts);
    const revealButton = screen.getAllByText('Révéler')[0];
    await fireEvent.click(revealButton);
    
    const cancelButton = screen.getByText('Annuler');
    await fireEvent.click(cancelButton);
    
    expect(screen.queryByText('Vérification')).not.toBeInTheDocument();
  });
});
