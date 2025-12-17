new EventSource('http://127.0.0.1:8088/esbuild').addEventListener('change', () => location.reload())

// node_modules/svelte/src/runtime/internal/utils.js
function noop() {
}
function run(fn) {
  return fn();
}
function blank_object() {
  return /* @__PURE__ */ Object.create(null);
}
function run_all(fns) {
  fns.forEach(run);
}
function is_function(thing) {
  return typeof thing === "function";
}
function safe_not_equal(a, b) {
  return a != a ? b == b : a !== b || a && typeof a === "object" || typeof a === "function";
}
function is_empty(obj) {
  return Object.keys(obj).length === 0;
}

// node_modules/svelte/src/runtime/internal/globals.js
var globals = typeof window !== "undefined" ? window : typeof globalThis !== "undefined" ? globalThis : (
  // @ts-ignore Node typings have this
  global
);

// node_modules/svelte/src/runtime/internal/ResizeObserverSingleton.js
var ResizeObserverSingleton = class _ResizeObserverSingleton {
  /**
   * @private
   * @readonly
   * @type {WeakMap<Element, import('./private.js').Listener>}
   */
  _listeners = "WeakMap" in globals ? /* @__PURE__ */ new WeakMap() : void 0;
  /**
   * @private
   * @type {ResizeObserver}
   */
  _observer = void 0;
  /** @type {ResizeObserverOptions} */
  options;
  /** @param {ResizeObserverOptions} options */
  constructor(options) {
    this.options = options;
  }
  /**
   * @param {Element} element
   * @param {import('./private.js').Listener} listener
   * @returns {() => void}
   */
  observe(element2, listener) {
    this._listeners.set(element2, listener);
    this._getObserver().observe(element2, this.options);
    return () => {
      this._listeners.delete(element2);
      this._observer.unobserve(element2);
    };
  }
  /**
   * @private
   */
  _getObserver() {
    return this._observer ?? (this._observer = new ResizeObserver((entries) => {
      for (const entry of entries) {
        _ResizeObserverSingleton.entries.set(entry.target, entry);
        this._listeners.get(entry.target)?.(entry);
      }
    }));
  }
};
ResizeObserverSingleton.entries = "WeakMap" in globals ? /* @__PURE__ */ new WeakMap() : void 0;

// node_modules/svelte/src/runtime/internal/dom.js
var is_hydrating = false;
function start_hydrating() {
  is_hydrating = true;
}
function end_hydrating() {
  is_hydrating = false;
}
function append(target, node) {
  target.appendChild(node);
}
function append_styles(target, style_sheet_id, styles) {
  const append_styles_to = get_root_for_style(target);
  if (!append_styles_to.getElementById(style_sheet_id)) {
    const style = element("style");
    style.id = style_sheet_id;
    style.textContent = styles;
    append_stylesheet(append_styles_to, style);
  }
}
function get_root_for_style(node) {
  if (!node)
    return document;
  const root = node.getRootNode ? node.getRootNode() : node.ownerDocument;
  if (root && /** @type {ShadowRoot} */
  root.host) {
    return (
      /** @type {ShadowRoot} */
      root
    );
  }
  return node.ownerDocument;
}
function append_stylesheet(node, style) {
  append(
    /** @type {Document} */
    node.head || node,
    style
  );
  return style.sheet;
}
function insert(target, node, anchor) {
  target.insertBefore(node, anchor || null);
}
function detach(node) {
  if (node.parentNode) {
    node.parentNode.removeChild(node);
  }
}
function element(name2) {
  return document.createElement(name2);
}
function text(data) {
  return document.createTextNode(data);
}
function space() {
  return text(" ");
}
function attr(node, attribute, value) {
  if (value == null)
    node.removeAttribute(attribute);
  else if (node.getAttribute(attribute) !== value)
    node.setAttribute(attribute, value);
}
function children(element2) {
  return Array.from(element2.childNodes);
}
function set_data(text2, data) {
  data = "" + data;
  if (text2.data === data)
    return;
  text2.data = /** @type {string} */
  data;
}
function toggle_class(element2, name2, toggle) {
  element2.classList.toggle(name2, !!toggle);
}
function get_custom_elements_slots(element2) {
  const result = {};
  element2.childNodes.forEach(
    /** @param {Element} node */
    (node) => {
      result[node.slot || "default"] = true;
    }
  );
  return result;
}

// node_modules/svelte/src/runtime/internal/lifecycle.js
var current_component;
function set_current_component(component) {
  current_component = component;
}
function get_current_component() {
  if (!current_component)
    throw new Error("Function called outside component initialization");
  return current_component;
}
function onMount(fn) {
  get_current_component().$$.on_mount.push(fn);
}

// node_modules/svelte/src/runtime/internal/scheduler.js
var dirty_components = [];
var binding_callbacks = [];
var render_callbacks = [];
var flush_callbacks = [];
var resolved_promise = /* @__PURE__ */ Promise.resolve();
var update_scheduled = false;
function schedule_update() {
  if (!update_scheduled) {
    update_scheduled = true;
    resolved_promise.then(flush);
  }
}
function add_render_callback(fn) {
  render_callbacks.push(fn);
}
var seen_callbacks = /* @__PURE__ */ new Set();
var flushidx = 0;
function flush() {
  if (flushidx !== 0) {
    return;
  }
  const saved_component = current_component;
  do {
    try {
      while (flushidx < dirty_components.length) {
        const component = dirty_components[flushidx];
        flushidx++;
        set_current_component(component);
        update(component.$$);
      }
    } catch (e) {
      dirty_components.length = 0;
      flushidx = 0;
      throw e;
    }
    set_current_component(null);
    dirty_components.length = 0;
    flushidx = 0;
    while (binding_callbacks.length)
      binding_callbacks.pop()();
    for (let i = 0; i < render_callbacks.length; i += 1) {
      const callback = render_callbacks[i];
      if (!seen_callbacks.has(callback)) {
        seen_callbacks.add(callback);
        callback();
      }
    }
    render_callbacks.length = 0;
  } while (dirty_components.length);
  while (flush_callbacks.length) {
    flush_callbacks.pop()();
  }
  update_scheduled = false;
  seen_callbacks.clear();
  set_current_component(saved_component);
}
function update($$) {
  if ($$.fragment !== null) {
    $$.update();
    run_all($$.before_update);
    const dirty = $$.dirty;
    $$.dirty = [-1];
    $$.fragment && $$.fragment.p($$.ctx, dirty);
    $$.after_update.forEach(add_render_callback);
  }
}
function flush_render_callbacks(fns) {
  const filtered = [];
  const targets = [];
  render_callbacks.forEach((c) => fns.indexOf(c) === -1 ? filtered.push(c) : targets.push(c));
  targets.forEach((c) => c());
  render_callbacks = filtered;
}

// node_modules/svelte/src/runtime/internal/transitions.js
var outroing = /* @__PURE__ */ new Set();
function transition_in(block, local) {
  if (block && block.i) {
    outroing.delete(block);
    block.i(local);
  }
}

// node_modules/svelte/src/shared/boolean_attributes.js
var _boolean_attributes = (
  /** @type {const} */
  [
    "allowfullscreen",
    "allowpaymentrequest",
    "async",
    "autofocus",
    "autoplay",
    "checked",
    "controls",
    "default",
    "defer",
    "disabled",
    "formnovalidate",
    "hidden",
    "inert",
    "ismap",
    "loop",
    "multiple",
    "muted",
    "nomodule",
    "novalidate",
    "open",
    "playsinline",
    "readonly",
    "required",
    "reversed",
    "selected"
  ]
);
var boolean_attributes = /* @__PURE__ */ new Set([..._boolean_attributes]);

// node_modules/svelte/src/runtime/internal/Component.js
function mount_component(component, target, anchor) {
  const { fragment, after_update } = component.$$;
  fragment && fragment.m(target, anchor);
  add_render_callback(() => {
    const new_on_destroy = component.$$.on_mount.map(run).filter(is_function);
    if (component.$$.on_destroy) {
      component.$$.on_destroy.push(...new_on_destroy);
    } else {
      run_all(new_on_destroy);
    }
    component.$$.on_mount = [];
  });
  after_update.forEach(add_render_callback);
}
function destroy_component(component, detaching) {
  const $$ = component.$$;
  if ($$.fragment !== null) {
    flush_render_callbacks($$.after_update);
    run_all($$.on_destroy);
    $$.fragment && $$.fragment.d(detaching);
    $$.on_destroy = $$.fragment = null;
    $$.ctx = [];
  }
}
function make_dirty(component, i) {
  if (component.$$.dirty[0] === -1) {
    dirty_components.push(component);
    schedule_update();
    component.$$.dirty.fill(0);
  }
  component.$$.dirty[i / 31 | 0] |= 1 << i % 31;
}
function init(component, options, instance2, create_fragment2, not_equal, props, append_styles2 = null, dirty = [-1]) {
  const parent_component = current_component;
  set_current_component(component);
  const $$ = component.$$ = {
    fragment: null,
    ctx: [],
    // state
    props,
    update: noop,
    not_equal,
    bound: blank_object(),
    // lifecycle
    on_mount: [],
    on_destroy: [],
    on_disconnect: [],
    before_update: [],
    after_update: [],
    context: new Map(options.context || (parent_component ? parent_component.$$.context : [])),
    // everything else
    callbacks: blank_object(),
    dirty,
    skip_bound: false,
    root: options.target || parent_component.$$.root
  };
  append_styles2 && append_styles2($$.root);
  let ready = false;
  $$.ctx = instance2 ? instance2(component, options.props || {}, (i, ret, ...rest) => {
    const value = rest.length ? rest[0] : ret;
    if ($$.ctx && not_equal($$.ctx[i], $$.ctx[i] = value)) {
      if (!$$.skip_bound && $$.bound[i])
        $$.bound[i](value);
      if (ready)
        make_dirty(component, i);
    }
    return ret;
  }) : [];
  $$.update();
  ready = true;
  run_all($$.before_update);
  $$.fragment = create_fragment2 ? create_fragment2($$.ctx) : false;
  if (options.target) {
    if (options.hydrate) {
      start_hydrating();
      const nodes = children(options.target);
      $$.fragment && $$.fragment.l(nodes);
      nodes.forEach(detach);
    } else {
      $$.fragment && $$.fragment.c();
    }
    if (options.intro)
      transition_in(component.$$.fragment);
    mount_component(component, options.target, options.anchor);
    end_hydrating();
    flush();
  }
  set_current_component(parent_component);
}
var SvelteElement;
if (typeof HTMLElement === "function") {
  SvelteElement = class extends HTMLElement {
    /** The Svelte component constructor */
    $$ctor;
    /** Slots */
    $$s;
    /** The Svelte component instance */
    $$c;
    /** Whether or not the custom element is connected */
    $$cn = false;
    /** Component props data */
    $$d = {};
    /** `true` if currently in the process of reflecting component props back to attributes */
    $$r = false;
    /** @type {Record<string, CustomElementPropDefinition>} Props definition (name, reflected, type etc) */
    $$p_d = {};
    /** @type {Record<string, Function[]>} Event listeners */
    $$l = {};
    /** @type {Map<Function, Function>} Event listener unsubscribe functions */
    $$l_u = /* @__PURE__ */ new Map();
    constructor($$componentCtor, $$slots, use_shadow_dom) {
      super();
      this.$$ctor = $$componentCtor;
      this.$$s = $$slots;
      if (use_shadow_dom) {
        this.attachShadow({ mode: "open" });
      }
    }
    addEventListener(type, listener, options) {
      this.$$l[type] = this.$$l[type] || [];
      this.$$l[type].push(listener);
      if (this.$$c) {
        const unsub = this.$$c.$on(type, listener);
        this.$$l_u.set(listener, unsub);
      }
      super.addEventListener(type, listener, options);
    }
    removeEventListener(type, listener, options) {
      super.removeEventListener(type, listener, options);
      if (this.$$c) {
        const unsub = this.$$l_u.get(listener);
        if (unsub) {
          unsub();
          this.$$l_u.delete(listener);
        }
      }
      if (this.$$l[type]) {
        const idx = this.$$l[type].indexOf(listener);
        if (idx >= 0) {
          this.$$l[type].splice(idx, 1);
        }
      }
    }
    async connectedCallback() {
      this.$$cn = true;
      if (!this.$$c) {
        let create_slot = function(name2) {
          return () => {
            let node;
            const obj = {
              c: function create() {
                node = element("slot");
                if (name2 !== "default") {
                  attr(node, "name", name2);
                }
              },
              /**
               * @param {HTMLElement} target
               * @param {HTMLElement} [anchor]
               */
              m: function mount(target, anchor) {
                insert(target, node, anchor);
              },
              d: function destroy(detaching) {
                if (detaching) {
                  detach(node);
                }
              }
            };
            return obj;
          };
        };
        await Promise.resolve();
        if (!this.$$cn || this.$$c) {
          return;
        }
        const $$slots = {};
        const existing_slots = get_custom_elements_slots(this);
        for (const name2 of this.$$s) {
          if (name2 in existing_slots) {
            $$slots[name2] = [create_slot(name2)];
          }
        }
        for (const attribute of this.attributes) {
          const name2 = this.$$g_p(attribute.name);
          if (!(name2 in this.$$d)) {
            this.$$d[name2] = get_custom_element_value(name2, attribute.value, this.$$p_d, "toProp");
          }
        }
        for (const key in this.$$p_d) {
          if (!(key in this.$$d) && this[key] !== void 0) {
            this.$$d[key] = this[key];
            delete this[key];
          }
        }
        this.$$c = new this.$$ctor({
          target: this.shadowRoot || this,
          props: {
            ...this.$$d,
            $$slots,
            $$scope: {
              ctx: []
            }
          }
        });
        const reflect_attributes = () => {
          this.$$r = true;
          for (const key in this.$$p_d) {
            this.$$d[key] = this.$$c.$$.ctx[this.$$c.$$.props[key]];
            if (this.$$p_d[key].reflect) {
              const attribute_value = get_custom_element_value(
                key,
                this.$$d[key],
                this.$$p_d,
                "toAttribute"
              );
              if (attribute_value == null) {
                this.removeAttribute(this.$$p_d[key].attribute || key);
              } else {
                this.setAttribute(this.$$p_d[key].attribute || key, attribute_value);
              }
            }
          }
          this.$$r = false;
        };
        this.$$c.$$.after_update.push(reflect_attributes);
        reflect_attributes();
        for (const type in this.$$l) {
          for (const listener of this.$$l[type]) {
            const unsub = this.$$c.$on(type, listener);
            this.$$l_u.set(listener, unsub);
          }
        }
        this.$$l = {};
      }
    }
    // We don't need this when working within Svelte code, but for compatibility of people using this outside of Svelte
    // and setting attributes through setAttribute etc, this is helpful
    attributeChangedCallback(attr2, _oldValue, newValue) {
      if (this.$$r)
        return;
      attr2 = this.$$g_p(attr2);
      this.$$d[attr2] = get_custom_element_value(attr2, newValue, this.$$p_d, "toProp");
      this.$$c?.$set({ [attr2]: this.$$d[attr2] });
    }
    disconnectedCallback() {
      this.$$cn = false;
      Promise.resolve().then(() => {
        if (!this.$$cn && this.$$c) {
          this.$$c.$destroy();
          this.$$c = void 0;
        }
      });
    }
    $$g_p(attribute_name) {
      return Object.keys(this.$$p_d).find(
        (key) => this.$$p_d[key].attribute === attribute_name || !this.$$p_d[key].attribute && key.toLowerCase() === attribute_name
      ) || attribute_name;
    }
  };
}
function get_custom_element_value(prop, value, props_definition, transform) {
  const type = props_definition[prop]?.type;
  value = type === "Boolean" && typeof value !== "boolean" ? value != null : value;
  if (!transform || !props_definition[prop]) {
    return value;
  } else if (transform === "toAttribute") {
    switch (type) {
      case "Object":
      case "Array":
        return value == null ? null : JSON.stringify(value);
      case "Boolean":
        return value ? "" : null;
      case "Number":
        return value == null ? null : value;
      default:
        return value;
    }
  } else {
    switch (type) {
      case "Object":
      case "Array":
        return value && JSON.parse(value);
      case "Boolean":
        return value;
      case "Number":
        return value != null ? +value : value;
      default:
        return value;
    }
  }
}
function create_custom_element(Component, props_definition, slots, accessors, use_shadow_dom, extend) {
  let Class = class extends SvelteElement {
    constructor() {
      super(Component, slots, use_shadow_dom);
      this.$$p_d = props_definition;
    }
    static get observedAttributes() {
      return Object.keys(props_definition).map(
        (key) => (props_definition[key].attribute || key).toLowerCase()
      );
    }
  };
  Object.keys(props_definition).forEach((prop) => {
    Object.defineProperty(Class.prototype, prop, {
      get() {
        return this.$$c && prop in this.$$c ? this.$$c[prop] : this.$$d[prop];
      },
      set(value) {
        value = get_custom_element_value(prop, value, props_definition);
        this.$$d[prop] = value;
        this.$$c?.$set({ [prop]: value });
      }
    });
  });
  accessors.forEach((accessor) => {
    Object.defineProperty(Class.prototype, accessor, {
      get() {
        return this.$$c?.[accessor];
      }
    });
  });
  if (extend) {
    Class = extend(Class);
  }
  Component.element = /** @type {any} */
  Class;
  return Class;
}
var SvelteComponent = class {
  /**
   * ### PRIVATE API
   *
   * Do not use, may change at any time
   *
   * @type {any}
   */
  $$ = void 0;
  /**
   * ### PRIVATE API
   *
   * Do not use, may change at any time
   *
   * @type {any}
   */
  $$set = void 0;
  /** @returns {void} */
  $destroy() {
    destroy_component(this, 1);
    this.$destroy = noop;
  }
  /**
   * @template {Extract<keyof Events, string>} K
   * @param {K} type
   * @param {((e: Events[K]) => void) | null | undefined} callback
   * @returns {() => void}
   */
  $on(type, callback) {
    if (!is_function(callback)) {
      return noop;
    }
    const callbacks = this.$$.callbacks[type] || (this.$$.callbacks[type] = []);
    callbacks.push(callback);
    return () => {
      const index = callbacks.indexOf(callback);
      if (index !== -1)
        callbacks.splice(index, 1);
    };
  }
  /**
   * @param {Partial<Props>} props
   * @returns {void}
   */
  $set(props) {
    if (this.$$set && !is_empty(props)) {
      this.$$.skip_bound = true;
      this.$$set(props);
      this.$$.skip_bound = false;
    }
  }
};

// node_modules/svelte/src/shared/version.js
var PUBLIC_VERSION = "4";

// node_modules/svelte/src/runtime/internal/disclose-version/index.js
if (typeof window !== "undefined")
  (window.__svelte || (window.__svelte = { v: /* @__PURE__ */ new Set() })).v.add(PUBLIC_VERSION);

// ihm/app.svelte
function add_css(target) {
  append_styles(target, "svelte-1rual6n", `@import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css");@import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");.banner.svelte-1rual6n.svelte-1rual6n{position:relative;height:calc(100vh - var(--header-height));display:flex;flex-direction:column;overflow:hidden;background:url('../assets/background.webp') center/cover no-repeat}.banner.svelte-1rual6n.svelte-1rual6n::before{content:'';position:absolute;top:0;left:0;right:0;bottom:0;background-image:inherit;background-size:cover;background-position:center;z-index:0}.banner-content.svelte-1rual6n.svelte-1rual6n{position:relative;z-index:1;color:white;text-align:left;text-shadow:2px 2px 4px rgba(0, 0, 0, 0.9);min-width:fit-content;padding-left:20px;margin:auto auto auto 0}.banner-content.svelte-1rual6n h1.svelte-1rual6n,.banner-content.svelte-1rual6n h2.svelte-1rual6n{margin:0;padding:0.2em 0;text-align:left;transition:opacity 0.5s ease-in-out}.banner-content.svelte-1rual6n h2.fade-out.svelte-1rual6n{opacity:0}.parcours.svelte-1rual6n.svelte-1rual6n{align-self:center;margin-top:auto;margin-bottom:0.5em;z-index:1;color:white;font-size:1.5rem;text-shadow:2px 2px 4px rgba(0, 0, 0, 0.9)}.scroll-indicator.svelte-1rual6n.svelte-1rual6n{align-self:center;margin-bottom:1rem;font-size:1.5rem;color:white;text-shadow:2px 2px 4px rgba(0, 0, 0, 0.9);z-index:1;animation:svelte-1rual6n-bounce 2s infinite}@keyframes svelte-1rual6n-bounce{0%,20%,50%,80%,100%{transform:translateY(0)}40%{transform:translateY(-10px)}60%{transform:translateY(-5px)}}.cv-container.svelte-1rual6n.svelte-1rual6n{display:flex;flex-wrap:wrap;max-width:1200px;margin:2rem auto;overflow:hidden}.cv-container.svelte-1rual6n h2.svelte-1rual6n,.cv-container.svelte-1rual6n h3.svelte-1rual6n{padding-bottom:0.5rem;margin-bottom:1rem}.cv-container.svelte-1rual6n h2.svelte-1rual6n{font-size:1.8rem}.cv-container.svelte-1rual6n h3.svelte-1rual6n{font-size:1.2rem;border-bottom:none;margin-top:1.5rem}.cv-container.svelte-1rual6n section.svelte-1rual6n{margin-bottom:2rem}.cv-container.svelte-1rual6n ul.svelte-1rual6n{padding-left:20px}.cv-container.svelte-1rual6n li.svelte-1rual6n{margin-bottom:0.5rem}.cv-container.svelte-1rual6n article.svelte-1rual6n{margin-bottom:1.5rem}@media(max-width: 768px){.cv-container.svelte-1rual6n.svelte-1rual6n{flex-direction:column}}`);
}
function create_fragment(ctx) {
  let div2;
  let div0;
  let h1;
  let t1;
  let h20;
  let t2;
  let t3;
  let span;
  let t5;
  let div1;
  let t6;
  let div3;
  return {
    c() {
      div2 = element("div");
      div0 = element("div");
      h1 = element("h1");
      h1.textContent = `${name}`;
      t1 = space();
      h20 = element("h2");
      t2 = text(
        /*currentH2Text*/
        ctx[0]
      );
      t3 = space();
      span = element("span");
      span.textContent = "Mon parcours";
      t5 = space();
      div1 = element("div");
      div1.innerHTML = `<i class="bi bi-arrow-down-circle"></i>`;
      t6 = space();
      div3 = element("div");
      div3.innerHTML = `<section class="experiences svelte-1rual6n"><h2 class="svelte-1rual6n">Exp\xE9riences professionnelles</h2> <article class="svelte-1rual6n"><h3 class="svelte-1rual6n">D\xE9veloppeur fullstack [Alternance]</h3> <p><strong>Softinnov | mars 2024 \xE0 mars 2026</strong></p> <ul class="svelte-1rual6n"><li class="svelte-1rual6n">Conception et d\xE9veloppement d&#39;applications et de solutions web pour divers clients.</li> <li class="svelte-1rual6n">Administration de syst\xE8mes, migration de base de donn\xE9es.</li></ul></article> <article class="svelte-1rual6n"><h3 class="svelte-1rual6n">Etude du gaspillage alimentaire [Service Civique]</h3> <p><strong>Mairie de Saint-Jean de V\xE9das | novembre 2022 \xE0 juillet 2023</strong></p> <ul class="svelte-1rual6n"><li class="svelte-1rual6n">Diagnostique du gaspillage alimentaire \xE0 travers la collecte et l&#39;\xE9tude de donn\xE9es.</li> <li class="svelte-1rual6n">R\xE9daction d&#39;un rapport et pr\xE9sentation des r\xE9sultats aux \xE9lus locaux.</li> <li class="svelte-1rual6n">Cr\xE9ation d&#39;un jeu vid\xE9o \xE9ducatif en javascript pour sensibiliser aux \xE9cogestes.</li></ul></article> <article class="svelte-1rual6n"><h3 class="svelte-1rual6n">Support sur un projet d&#39;innovation [Stage]</h3> <p><strong>IDEMIA \u2013 R&amp;D Sophia-Antipolis | mai 2022 \xE0 ao\xFBt 2022</strong></p> <ul class="svelte-1rual6n"><li class="svelte-1rual6n">Cr\xE9ation et automatisation de tests d&#39;UI.</li> <li class="svelte-1rual6n">Utilisation d&#39;un framework de test \xAB End to End \xBB (Python, Selenium, Serenity, Behave).</li></ul></article></section> <section class="formations svelte-1rual6n"><h2 class="svelte-1rual6n">Formations</h2> <article class="svelte-1rual6n"><h3 class="svelte-1rual6n">Mast\xE8re Expert en Ing\xE9nierie Logicielle</h3> <p><strong>ISCOD | 2023 - 2025</strong></p></article> <article class="svelte-1rual6n"><h3 class="svelte-1rual6n">Licence de math\xE9matiques et informatique appliqu\xE9es</h3> <p><strong>Paul Val\xE9ry Montpellier 3 | 2019 - 2022</strong></p></article> <article class="svelte-1rual6n"><h3 class="svelte-1rual6n">Baccalaur\xE9at Scientifique</h3> <p><strong>Lyc\xE9e Laetitia Bonaparte | 2013 - 2016</strong></p></article></section>`;
      attr(h1, "class", "h1 svelte-1rual6n");
      attr(h20, "class", "h2 svelte-1rual6n");
      toggle_class(
        h20,
        "fade-out",
        /*isFading*/
        ctx[1]
      );
      attr(div0, "class", "banner-content svelte-1rual6n");
      attr(span, "class", "parcours svelte-1rual6n");
      attr(div1, "class", "scroll-indicator svelte-1rual6n");
      attr(div2, "class", "banner svelte-1rual6n");
      attr(div3, "class", "cv-container svelte-1rual6n");
    },
    m(target, anchor) {
      insert(target, div2, anchor);
      append(div2, div0);
      append(div0, h1);
      append(div0, t1);
      append(div0, h20);
      append(h20, t2);
      append(div2, t3);
      append(div2, span);
      append(div2, t5);
      append(div2, div1);
      insert(target, t6, anchor);
      insert(target, div3, anchor);
    },
    p(ctx2, [dirty]) {
      if (dirty & /*currentH2Text*/
      1)
        set_data(
          t2,
          /*currentH2Text*/
          ctx2[0]
        );
      if (dirty & /*isFading*/
      2) {
        toggle_class(
          h20,
          "fade-out",
          /*isFading*/
          ctx2[1]
        );
      }
    },
    i: noop,
    o: noop,
    d(detaching) {
      if (detaching) {
        detach(div2);
        detach(t6);
        detach(div3);
      }
    }
  };
}
var name = "Cl\xE9ment Calia";
var inge = "Expert en ing\xE9nierie logicielle";
var fullstack = "D\xE9veloppeur Fullstack";
var transitionDuration = 500;
var displayDuration = 3e3;
function instance($$self, $$props, $$invalidate) {
  let currentH2Text = inge;
  let isFading = false;
  onMount(() => {
    const interval = setInterval(
      () => {
        $$invalidate(1, isFading = true);
        setTimeout(
          () => {
            $$invalidate(0, currentH2Text = currentH2Text === inge ? fullstack : inge);
            $$invalidate(1, isFading = false);
          },
          transitionDuration
        );
      },
      displayDuration + transitionDuration
    );
    return () => clearInterval(interval);
  });
  return [currentH2Text, isFading];
}
var App = class extends SvelteComponent {
  constructor(options) {
    super();
    init(this, options, instance, create_fragment, safe_not_equal, {}, add_css);
  }
};
customElements.define("index-portfolio", create_custom_element(App, {}, [], [], true));
var app_default = App;
export {
  app_default as default
};
