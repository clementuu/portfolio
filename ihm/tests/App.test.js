import { render, screen } from '@testing-library/svelte';
import { describe, it, expect } from 'vitest';
import App from '../App.svelte';

describe('App - static UI checks', () => {
  it('renders the name "Clément Calia" inside the h1 strong', () => {
    const { container } = render(App);
    const strong = container.querySelector('h1.name strong');
    expect(strong).toBeTruthy();
    expect(strong?.textContent).toBe('Clément Calia');
  });

  it('renders the initial job title inside the h2 strong', () => {
    const { container } = render(App);
    const strong = container.querySelector('h2.metier strong');
    expect(strong).toBeTruthy();
    expect(strong?.textContent).toBe('Expert en ingénierie logicielle');
  });

  it('does not apply fade-out class initially on the h2', () => {
    const { container } = render(App);
    const h2 = container.querySelector('h2.metier');
    expect(h2).toBeTruthy();
    expect(h2?.classList.contains('fade-out')).toBe(false);
  });

  it('contains the parcours label', () => {
    render(App);
    expect(screen.getByText('Mon parcours')).toBeInTheDocument();
  });

  it('renders the profile image with class my-pic and an img element', () => {
    const { container } = render(App);
    const img = container.querySelector('img.my-pic');
    expect(img).toBeTruthy();
    // vérifie qu'il y a bien une balise img (alt peut être vide)
    expect(img?.tagName.toLowerCase()).toBe('img');
  });

  it('renders the scroll indicator icon element', () => {
    const { container } = render(App);
    // on vérifie la présence de l'icône via ses classes
    const icon = container.querySelector('.scroll-indicator .bi.bi-arrow-down-circle');
    expect(icon).toBeTruthy();
  });

  it('has a root element with class banner', () => {
    const { container } = render(App);
    const banner = container.querySelector('.banner');
    expect(banner).toBeTruthy();
  });
});
