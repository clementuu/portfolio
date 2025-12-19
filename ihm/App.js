new EventSource('http://127.0.0.1:8088/esbuild').addEventListener('change', () => location.reload())

// node_modules/svelte/src/runtime/internal/utils.js
function noop() {
}
function assign(tar, src) {
  for (const k in src)
    tar[k] = src[k];
  return (
    /** @type {T & S} */
    tar
  );
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
var src_url_equal_anchor;
function src_url_equal(element_src, url) {
  if (element_src === url)
    return true;
  if (!src_url_equal_anchor) {
    src_url_equal_anchor = document.createElement("a");
  }
  src_url_equal_anchor.href = url;
  return element_src === src_url_equal_anchor.href;
}
function is_empty(obj) {
  return Object.keys(obj).length === 0;
}
function create_slot(definition, ctx, $$scope, fn) {
  if (definition) {
    const slot_ctx = get_slot_context(definition, ctx, $$scope, fn);
    return definition[0](slot_ctx);
  }
}
function get_slot_context(definition, ctx, $$scope, fn) {
  return definition[1] && fn ? assign($$scope.ctx.slice(), definition[1](fn(ctx))) : $$scope.ctx;
}
function get_slot_changes(definition, $$scope, dirty, fn) {
  if (definition[2] && fn) {
    const lets = definition[2](fn(dirty));
    if ($$scope.dirty === void 0) {
      return lets;
    }
    if (typeof lets === "object") {
      const merged = [];
      const len = Math.max($$scope.dirty.length, lets.length);
      for (let i = 0; i < len; i += 1) {
        merged[i] = $$scope.dirty[i] | lets[i];
      }
      return merged;
    }
    return $$scope.dirty | lets;
  }
  return $$scope.dirty;
}
function update_slot_base(slot, slot_definition, ctx, $$scope, slot_changes, get_slot_context_fn) {
  if (slot_changes) {
    const slot_context = get_slot_context(slot_definition, ctx, $$scope, get_slot_context_fn);
    slot.p(slot_context, slot_changes);
  }
}
function get_all_dirty_from_scope($$scope) {
  if ($$scope.ctx.length > 32) {
    const dirty = [];
    const length = $$scope.ctx.length / 32;
    for (let i = 0; i < length; i++) {
      dirty[i] = -1;
    }
    return dirty;
  }
  return -1;
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
function destroy_each(iterations, detaching) {
  for (let i = 0; i < iterations.length; i += 1) {
    if (iterations[i])
      iterations[i].d(detaching);
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
function empty() {
  return text("");
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
function set_style(node, key, value, important) {
  if (value == null) {
    node.style.removeProperty(key);
  } else {
    node.style.setProperty(key, value, important ? "important" : "");
  }
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
var outros;
function group_outros() {
  outros = {
    r: 0,
    c: [],
    p: outros
    // parent group
  };
}
function check_outros() {
  if (!outros.r) {
    run_all(outros.c);
  }
  outros = outros.p;
}
function transition_in(block, local) {
  if (block && block.i) {
    outroing.delete(block);
    block.i(local);
  }
}
function transition_out(block, local, detach2, callback) {
  if (block && block.o) {
    if (outroing.has(block))
      return;
    outroing.add(block);
    outros.c.push(() => {
      outroing.delete(block);
      if (callback) {
        if (detach2)
          block.d(1);
        callback();
      }
    });
    block.o(local);
  } else if (callback) {
    callback();
  }
}

// node_modules/svelte/src/runtime/internal/each.js
function ensure_array_like(array_like_or_iterator) {
  return array_like_or_iterator?.length !== void 0 ? array_like_or_iterator : Array.from(array_like_or_iterator);
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
function create_component(block) {
  block && block.c();
}
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
function init(component, options, instance7, create_fragment7, not_equal, props, append_styles2 = null, dirty = [-1]) {
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
  $$.ctx = instance7 ? instance7(component, options.props || {}, (i, ret, ...rest) => {
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
  $$.fragment = create_fragment7 ? create_fragment7($$.ctx) : false;
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
        let create_slot2 = function(name2) {
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
            $$slots[name2] = [create_slot2(name2)];
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

// ihm/cv/card.svelte
function add_css(target) {
  append_styles(target, "svelte-lmmzta", "article.svelte-lmmzta{background-color:#f9f9f9;padding:1.5rem;margin-bottom:1.5rem;border-radius:8px;box-shadow:0 4px 6px rgba(0, 0, 0, 0.1);transition:box-shadow 0.2s ease-in-out}h3.svelte-lmmzta{font-size:1.3rem;color:#333;margin-top:0;margin-bottom:0.5rem;border-bottom:none}p.svelte-lmmzta{font-size:1rem;color:#555;line-height:1.5}strong.svelte-lmmzta{color:#007bff}article ul{padding-left:20px;margin-top:1rem;color:#666}article li{margin-bottom:0.5rem}");
}
var get_subtitle_slot_changes = (dirty) => ({});
var get_subtitle_slot_context = (ctx) => ({});
var get_title_slot_changes = (dirty) => ({});
var get_title_slot_context = (ctx) => ({});
function create_fragment(ctx) {
  let article;
  let h3;
  let t0;
  let p;
  let strong;
  let t1;
  let current;
  const title_slot_template = (
    /*#slots*/
    ctx[1].title
  );
  const title_slot = create_slot(
    title_slot_template,
    ctx,
    /*$$scope*/
    ctx[0],
    get_title_slot_context
  );
  const subtitle_slot_template = (
    /*#slots*/
    ctx[1].subtitle
  );
  const subtitle_slot = create_slot(
    subtitle_slot_template,
    ctx,
    /*$$scope*/
    ctx[0],
    get_subtitle_slot_context
  );
  const default_slot_template = (
    /*#slots*/
    ctx[1].default
  );
  const default_slot = create_slot(
    default_slot_template,
    ctx,
    /*$$scope*/
    ctx[0],
    null
  );
  return {
    c() {
      article = element("article");
      h3 = element("h3");
      if (title_slot)
        title_slot.c();
      t0 = space();
      p = element("p");
      strong = element("strong");
      if (subtitle_slot)
        subtitle_slot.c();
      t1 = space();
      if (default_slot)
        default_slot.c();
      attr(h3, "class", "svelte-lmmzta");
      attr(strong, "class", "svelte-lmmzta");
      attr(p, "class", "svelte-lmmzta");
      attr(article, "class", "svelte-lmmzta");
    },
    m(target, anchor) {
      insert(target, article, anchor);
      append(article, h3);
      if (title_slot) {
        title_slot.m(h3, null);
      }
      append(article, t0);
      append(article, p);
      append(p, strong);
      if (subtitle_slot) {
        subtitle_slot.m(strong, null);
      }
      append(article, t1);
      if (default_slot) {
        default_slot.m(article, null);
      }
      current = true;
    },
    p(ctx2, [dirty]) {
      if (title_slot) {
        if (title_slot.p && (!current || dirty & /*$$scope*/
        1)) {
          update_slot_base(
            title_slot,
            title_slot_template,
            ctx2,
            /*$$scope*/
            ctx2[0],
            !current ? get_all_dirty_from_scope(
              /*$$scope*/
              ctx2[0]
            ) : get_slot_changes(
              title_slot_template,
              /*$$scope*/
              ctx2[0],
              dirty,
              get_title_slot_changes
            ),
            get_title_slot_context
          );
        }
      }
      if (subtitle_slot) {
        if (subtitle_slot.p && (!current || dirty & /*$$scope*/
        1)) {
          update_slot_base(
            subtitle_slot,
            subtitle_slot_template,
            ctx2,
            /*$$scope*/
            ctx2[0],
            !current ? get_all_dirty_from_scope(
              /*$$scope*/
              ctx2[0]
            ) : get_slot_changes(
              subtitle_slot_template,
              /*$$scope*/
              ctx2[0],
              dirty,
              get_subtitle_slot_changes
            ),
            get_subtitle_slot_context
          );
        }
      }
      if (default_slot) {
        if (default_slot.p && (!current || dirty & /*$$scope*/
        1)) {
          update_slot_base(
            default_slot,
            default_slot_template,
            ctx2,
            /*$$scope*/
            ctx2[0],
            !current ? get_all_dirty_from_scope(
              /*$$scope*/
              ctx2[0]
            ) : get_slot_changes(
              default_slot_template,
              /*$$scope*/
              ctx2[0],
              dirty,
              null
            ),
            null
          );
        }
      }
    },
    i(local) {
      if (current)
        return;
      transition_in(title_slot, local);
      transition_in(subtitle_slot, local);
      transition_in(default_slot, local);
      current = true;
    },
    o(local) {
      transition_out(title_slot, local);
      transition_out(subtitle_slot, local);
      transition_out(default_slot, local);
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(article);
      }
      if (title_slot)
        title_slot.d(detaching);
      if (subtitle_slot)
        subtitle_slot.d(detaching);
      if (default_slot)
        default_slot.d(detaching);
    }
  };
}
function instance($$self, $$props, $$invalidate) {
  let { $$slots: slots = {}, $$scope } = $$props;
  $$self.$$set = ($$props2) => {
    if ("$$scope" in $$props2)
      $$invalidate(0, $$scope = $$props2.$$scope);
  };
  return [$$scope, slots];
}
var Card = class extends SvelteComponent {
  constructor(options) {
    super();
    init(this, options, instance, create_fragment, safe_not_equal, {}, add_css);
  }
};
create_custom_element(Card, {}, ["title", "subtitle", "default"], [], true);
var card_default = Card;

// ihm/cv/formationCard.svelte
function create_title_slot(ctx) {
  let span;
  let t_value = (
    /*formation*/
    ctx[0].Intitule + ""
  );
  let t;
  return {
    c() {
      span = element("span");
      t = text(t_value);
      attr(span, "slot", "title");
    },
    m(target, anchor) {
      insert(target, span, anchor);
      append(span, t);
    },
    p(ctx2, dirty) {
      if (dirty & /*formation*/
      1 && t_value !== (t_value = /*formation*/
      ctx2[0].Intitule + ""))
        set_data(t, t_value);
    },
    d(detaching) {
      if (detaching) {
        detach(span);
      }
    }
  };
}
function create_subtitle_slot(ctx) {
  let span;
  let t0_value = (
    /*formation*/
    ctx[0].Etablissement + ""
  );
  let t0;
  let t1;
  let t2_value = (
    /*formation*/
    ctx[0].Periode + ""
  );
  let t2;
  return {
    c() {
      span = element("span");
      t0 = text(t0_value);
      t1 = text(" | ");
      t2 = text(t2_value);
      attr(span, "slot", "subtitle");
    },
    m(target, anchor) {
      insert(target, span, anchor);
      append(span, t0);
      append(span, t1);
      append(span, t2);
    },
    p(ctx2, dirty) {
      if (dirty & /*formation*/
      1 && t0_value !== (t0_value = /*formation*/
      ctx2[0].Etablissement + ""))
        set_data(t0, t0_value);
      if (dirty & /*formation*/
      1 && t2_value !== (t2_value = /*formation*/
      ctx2[0].Periode + ""))
        set_data(t2, t2_value);
    },
    d(detaching) {
      if (detaching) {
        detach(span);
      }
    }
  };
}
function create_fragment2(ctx) {
  let card;
  let current;
  card = new card_default({
    props: {
      $$slots: {
        subtitle: [create_subtitle_slot],
        title: [create_title_slot]
      },
      $$scope: { ctx }
    }
  });
  return {
    c() {
      create_component(card.$$.fragment);
    },
    m(target, anchor) {
      mount_component(card, target, anchor);
      current = true;
    },
    p(ctx2, [dirty]) {
      const card_changes = {};
      if (dirty & /*$$scope, formation*/
      3) {
        card_changes.$$scope = { dirty, ctx: ctx2 };
      }
      card.$set(card_changes);
    },
    i(local) {
      if (current)
        return;
      transition_in(card.$$.fragment, local);
      current = true;
    },
    o(local) {
      transition_out(card.$$.fragment, local);
      current = false;
    },
    d(detaching) {
      destroy_component(card, detaching);
    }
  };
}
function instance2($$self, $$props, $$invalidate) {
  let { formation } = $$props;
  $$self.$$set = ($$props2) => {
    if ("formation" in $$props2)
      $$invalidate(0, formation = $$props2.formation);
  };
  return [formation];
}
var FormationCard = class extends SvelteComponent {
  constructor(options) {
    super();
    init(this, options, instance2, create_fragment2, safe_not_equal, { formation: 0 });
  }
  get formation() {
    return this.$$.ctx[0];
  }
  set formation(formation) {
    this.$$set({ formation });
    flush();
  }
};
create_custom_element(FormationCard, { "formation": {} }, [], [], true);
var formationCard_default = FormationCard;

// ihm/cv/experienceCard.svelte
function get_each_context(ctx, list, i) {
  const child_ctx = ctx.slice();
  child_ctx[1] = list[i];
  return child_ctx;
}
function create_each_block(ctx) {
  let li;
  let t_value = (
    /*tache*/
    ctx[1] + ""
  );
  let t;
  return {
    c() {
      li = element("li");
      t = text(t_value);
    },
    m(target, anchor) {
      insert(target, li, anchor);
      append(li, t);
    },
    p(ctx2, dirty) {
      if (dirty & /*experience*/
      1 && t_value !== (t_value = /*tache*/
      ctx2[1] + ""))
        set_data(t, t_value);
    },
    d(detaching) {
      if (detaching) {
        detach(li);
      }
    }
  };
}
function create_default_slot(ctx) {
  let ul;
  let each_value = ensure_array_like(
    /*experience*/
    ctx[0].Taches
  );
  let each_blocks = [];
  for (let i = 0; i < each_value.length; i += 1) {
    each_blocks[i] = create_each_block(get_each_context(ctx, each_value, i));
  }
  return {
    c() {
      ul = element("ul");
      for (let i = 0; i < each_blocks.length; i += 1) {
        each_blocks[i].c();
      }
    },
    m(target, anchor) {
      insert(target, ul, anchor);
      for (let i = 0; i < each_blocks.length; i += 1) {
        if (each_blocks[i]) {
          each_blocks[i].m(ul, null);
        }
      }
    },
    p(ctx2, dirty) {
      if (dirty & /*experience*/
      1) {
        each_value = ensure_array_like(
          /*experience*/
          ctx2[0].Taches
        );
        let i;
        for (i = 0; i < each_value.length; i += 1) {
          const child_ctx = get_each_context(ctx2, each_value, i);
          if (each_blocks[i]) {
            each_blocks[i].p(child_ctx, dirty);
          } else {
            each_blocks[i] = create_each_block(child_ctx);
            each_blocks[i].c();
            each_blocks[i].m(ul, null);
          }
        }
        for (; i < each_blocks.length; i += 1) {
          each_blocks[i].d(1);
        }
        each_blocks.length = each_value.length;
      }
    },
    d(detaching) {
      if (detaching) {
        detach(ul);
      }
      destroy_each(each_blocks, detaching);
    }
  };
}
function create_title_slot2(ctx) {
  let span;
  let t0_value = (
    /*experience*/
    ctx[0].Intitule + ""
  );
  let t0;
  let t1;
  let t2_value = (
    /*experience*/
    ctx[0].Type + ""
  );
  let t2;
  let t3;
  return {
    c() {
      span = element("span");
      t0 = text(t0_value);
      t1 = text(" [");
      t2 = text(t2_value);
      t3 = text("]");
      attr(span, "slot", "title");
    },
    m(target, anchor) {
      insert(target, span, anchor);
      append(span, t0);
      append(span, t1);
      append(span, t2);
      append(span, t3);
    },
    p(ctx2, dirty) {
      if (dirty & /*experience*/
      1 && t0_value !== (t0_value = /*experience*/
      ctx2[0].Intitule + ""))
        set_data(t0, t0_value);
      if (dirty & /*experience*/
      1 && t2_value !== (t2_value = /*experience*/
      ctx2[0].Type + ""))
        set_data(t2, t2_value);
    },
    d(detaching) {
      if (detaching) {
        detach(span);
      }
    }
  };
}
function create_subtitle_slot2(ctx) {
  let span;
  let t0_value = (
    /*experience*/
    ctx[0].Structure + ""
  );
  let t0;
  let t1;
  let t2_value = (
    /*experience*/
    ctx[0].Periode + ""
  );
  let t2;
  return {
    c() {
      span = element("span");
      t0 = text(t0_value);
      t1 = text(" | ");
      t2 = text(t2_value);
      attr(span, "slot", "subtitle");
    },
    m(target, anchor) {
      insert(target, span, anchor);
      append(span, t0);
      append(span, t1);
      append(span, t2);
    },
    p(ctx2, dirty) {
      if (dirty & /*experience*/
      1 && t0_value !== (t0_value = /*experience*/
      ctx2[0].Structure + ""))
        set_data(t0, t0_value);
      if (dirty & /*experience*/
      1 && t2_value !== (t2_value = /*experience*/
      ctx2[0].Periode + ""))
        set_data(t2, t2_value);
    },
    d(detaching) {
      if (detaching) {
        detach(span);
      }
    }
  };
}
function create_fragment3(ctx) {
  let card;
  let current;
  card = new card_default({
    props: {
      $$slots: {
        subtitle: [create_subtitle_slot2],
        title: [create_title_slot2],
        default: [create_default_slot]
      },
      $$scope: { ctx }
    }
  });
  return {
    c() {
      create_component(card.$$.fragment);
    },
    m(target, anchor) {
      mount_component(card, target, anchor);
      current = true;
    },
    p(ctx2, [dirty]) {
      const card_changes = {};
      if (dirty & /*$$scope, experience*/
      17) {
        card_changes.$$scope = { dirty, ctx: ctx2 };
      }
      card.$set(card_changes);
    },
    i(local) {
      if (current)
        return;
      transition_in(card.$$.fragment, local);
      current = true;
    },
    o(local) {
      transition_out(card.$$.fragment, local);
      current = false;
    },
    d(detaching) {
      destroy_component(card, detaching);
    }
  };
}
function instance3($$self, $$props, $$invalidate) {
  let { experience } = $$props;
  $$self.$$set = ($$props2) => {
    if ("experience" in $$props2)
      $$invalidate(0, experience = $$props2.experience);
  };
  return [experience];
}
var ExperienceCard = class extends SvelteComponent {
  constructor(options) {
    super();
    init(this, options, instance3, create_fragment3, safe_not_equal, { experience: 0 });
  }
  get experience() {
    return this.$$.ctx[0];
  }
  set experience(experience) {
    this.$$set({ experience });
    flush();
  }
};
create_custom_element(ExperienceCard, { "experience": {} }, [], [], true);
var experienceCard_default = ExperienceCard;

// ihm/components/OnScrollAppear.svelte
function add_css2(target) {
  append_styles(target, "svelte-sjkedl", "div.svelte-sjkedl{opacity:0;transform:translateY(20px);transition:opacity 0.5s ease-out, transform 0.5s ease-out;display:block;width:100%}div.visible.svelte-sjkedl{opacity:1;transform:translateY(0)}");
}
function create_fragment4(ctx) {
  let div;
  let current;
  const default_slot_template = (
    /*#slots*/
    ctx[6].default
  );
  const default_slot = create_slot(
    default_slot_template,
    ctx,
    /*$$scope*/
    ctx[5],
    null
  );
  return {
    c() {
      div = element("div");
      if (default_slot)
        default_slot.c();
      set_style(
        div,
        "transition-delay",
        /*index*/
        ctx[0] * /*delay*/
        ctx[1] + "ms"
      );
      attr(div, "class", "svelte-sjkedl");
      toggle_class(
        div,
        "visible",
        /*visible*/
        ctx[2]
      );
    },
    m(target, anchor) {
      insert(target, div, anchor);
      if (default_slot) {
        default_slot.m(div, null);
      }
      ctx[7](div);
      current = true;
    },
    p(ctx2, [dirty]) {
      if (default_slot) {
        if (default_slot.p && (!current || dirty & /*$$scope*/
        32)) {
          update_slot_base(
            default_slot,
            default_slot_template,
            ctx2,
            /*$$scope*/
            ctx2[5],
            !current ? get_all_dirty_from_scope(
              /*$$scope*/
              ctx2[5]
            ) : get_slot_changes(
              default_slot_template,
              /*$$scope*/
              ctx2[5],
              dirty,
              null
            ),
            null
          );
        }
      }
      if (!current || dirty & /*index, delay*/
      3) {
        set_style(
          div,
          "transition-delay",
          /*index*/
          ctx2[0] * /*delay*/
          ctx2[1] + "ms"
        );
      }
      if (!current || dirty & /*visible*/
      4) {
        toggle_class(
          div,
          "visible",
          /*visible*/
          ctx2[2]
        );
      }
    },
    i(local) {
      if (current)
        return;
      transition_in(default_slot, local);
      current = true;
    },
    o(local) {
      transition_out(default_slot, local);
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(div);
      }
      if (default_slot)
        default_slot.d(detaching);
      ctx[7](null);
    }
  };
}
function instance4($$self, $$props, $$invalidate) {
  let { $$slots: slots = {}, $$scope } = $$props;
  let { index = 0 } = $$props;
  let { delay = 150 } = $$props;
  let { threshold = 0.1 } = $$props;
  let visible = false;
  let element2;
  onMount(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting) {
          $$invalidate(2, visible = true);
          observer.unobserve(element2);
        }
      },
      { threshold }
    );
    observer.observe(element2);
    return () => {
      if (element2)
        observer.unobserve(element2);
    };
  });
  function div_binding($$value) {
    binding_callbacks[$$value ? "unshift" : "push"](() => {
      element2 = $$value;
      $$invalidate(3, element2);
    });
  }
  $$self.$$set = ($$props2) => {
    if ("index" in $$props2)
      $$invalidate(0, index = $$props2.index);
    if ("delay" in $$props2)
      $$invalidate(1, delay = $$props2.delay);
    if ("threshold" in $$props2)
      $$invalidate(4, threshold = $$props2.threshold);
    if ("$$scope" in $$props2)
      $$invalidate(5, $$scope = $$props2.$$scope);
  };
  return [index, delay, visible, element2, threshold, $$scope, slots, div_binding];
}
var OnScrollAppear = class extends SvelteComponent {
  constructor(options) {
    super();
    init(this, options, instance4, create_fragment4, safe_not_equal, { index: 0, delay: 1, threshold: 4 }, add_css2);
  }
  get index() {
    return this.$$.ctx[0];
  }
  set index(index) {
    this.$$set({ index });
    flush();
  }
  get delay() {
    return this.$$.ctx[1];
  }
  set delay(delay) {
    this.$$set({ delay });
    flush();
  }
  get threshold() {
    return this.$$.ctx[4];
  }
  set threshold(threshold) {
    this.$$set({ threshold });
    flush();
  }
};
create_custom_element(OnScrollAppear, { "index": {}, "delay": {}, "threshold": {} }, ["default"], [], true);
var OnScrollAppear_default = OnScrollAppear;

// ihm/cv/cv.svelte
function add_css3(target) {
  append_styles(target, "svelte-1vov8fz", "section.svelte-1vov8fz{margin:2rem}h2.svelte-1vov8fz{font-size:2.5rem;font-weight:700;text-align:center;padding-bottom:0.5rem;margin-bottom:2rem}.error.svelte-1vov8fz{color:red;text-align:center}");
}
function get_each_context2(ctx, list, i) {
  const child_ctx = ctx.slice();
  child_ctx[5] = list[i];
  child_ctx[7] = i;
  return child_ctx;
}
function get_each_context_1(ctx, list, i) {
  const child_ctx = ctx.slice();
  child_ctx[8] = list[i];
  child_ctx[7] = i;
  return child_ctx;
}
function create_default_slot_3(ctx) {
  let h2;
  return {
    c() {
      h2 = element("h2");
      h2.textContent = "Formations";
      attr(h2, "class", "svelte-1vov8fz");
    },
    m(target, anchor) {
      insert(target, h2, anchor);
    },
    p: noop,
    d(detaching) {
      if (detaching) {
        detach(h2);
      }
    }
  };
}
function create_else_block_1(ctx) {
  let p;
  return {
    c() {
      p = element("p");
      p.textContent = "Chargement des formations...";
    },
    m(target, anchor) {
      insert(target, p, anchor);
    },
    p: noop,
    i: noop,
    o: noop,
    d(detaching) {
      if (detaching) {
        detach(p);
      }
    }
  };
}
function create_if_block_3(ctx) {
  let each_1_anchor;
  let current;
  let each_value_1 = ensure_array_like(
    /*formations*/
    ctx[0]
  );
  let each_blocks = [];
  for (let i = 0; i < each_value_1.length; i += 1) {
    each_blocks[i] = create_each_block_1(get_each_context_1(ctx, each_value_1, i));
  }
  const out = (i) => transition_out(each_blocks[i], 1, 1, () => {
    each_blocks[i] = null;
  });
  return {
    c() {
      for (let i = 0; i < each_blocks.length; i += 1) {
        each_blocks[i].c();
      }
      each_1_anchor = empty();
    },
    m(target, anchor) {
      for (let i = 0; i < each_blocks.length; i += 1) {
        if (each_blocks[i]) {
          each_blocks[i].m(target, anchor);
        }
      }
      insert(target, each_1_anchor, anchor);
      current = true;
    },
    p(ctx2, dirty) {
      if (dirty & /*formations*/
      1) {
        each_value_1 = ensure_array_like(
          /*formations*/
          ctx2[0]
        );
        let i;
        for (i = 0; i < each_value_1.length; i += 1) {
          const child_ctx = get_each_context_1(ctx2, each_value_1, i);
          if (each_blocks[i]) {
            each_blocks[i].p(child_ctx, dirty);
            transition_in(each_blocks[i], 1);
          } else {
            each_blocks[i] = create_each_block_1(child_ctx);
            each_blocks[i].c();
            transition_in(each_blocks[i], 1);
            each_blocks[i].m(each_1_anchor.parentNode, each_1_anchor);
          }
        }
        group_outros();
        for (i = each_value_1.length; i < each_blocks.length; i += 1) {
          out(i);
        }
        check_outros();
      }
    },
    i(local) {
      if (current)
        return;
      for (let i = 0; i < each_value_1.length; i += 1) {
        transition_in(each_blocks[i]);
      }
      current = true;
    },
    o(local) {
      each_blocks = each_blocks.filter(Boolean);
      for (let i = 0; i < each_blocks.length; i += 1) {
        transition_out(each_blocks[i]);
      }
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(each_1_anchor);
      }
      destroy_each(each_blocks, detaching);
    }
  };
}
function create_if_block_2(ctx) {
  let p;
  let t;
  return {
    c() {
      p = element("p");
      t = text(
        /*error*/
        ctx[1]
      );
      attr(p, "class", "error svelte-1vov8fz");
    },
    m(target, anchor) {
      insert(target, p, anchor);
      append(p, t);
    },
    p(ctx2, dirty) {
      if (dirty & /*error*/
      2)
        set_data(
          t,
          /*error*/
          ctx2[1]
        );
    },
    i: noop,
    o: noop,
    d(detaching) {
      if (detaching) {
        detach(p);
      }
    }
  };
}
function create_default_slot_2(ctx) {
  let formationcard;
  let t;
  let current;
  formationcard = new formationCard_default({
    props: { formation: (
      /*formation*/
      ctx[8]
    ) }
  });
  return {
    c() {
      create_component(formationcard.$$.fragment);
      t = space();
    },
    m(target, anchor) {
      mount_component(formationcard, target, anchor);
      insert(target, t, anchor);
      current = true;
    },
    p(ctx2, dirty) {
      const formationcard_changes = {};
      if (dirty & /*formations*/
      1)
        formationcard_changes.formation = /*formation*/
        ctx2[8];
      formationcard.$set(formationcard_changes);
    },
    i(local) {
      if (current)
        return;
      transition_in(formationcard.$$.fragment, local);
      current = true;
    },
    o(local) {
      transition_out(formationcard.$$.fragment, local);
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(t);
      }
      destroy_component(formationcard, detaching);
    }
  };
}
function create_each_block_1(ctx) {
  let onscrollappear;
  let current;
  onscrollappear = new OnScrollAppear_default({
    props: {
      index: (
        /*index*/
        ctx[7]
      ),
      $$slots: { default: [create_default_slot_2] },
      $$scope: { ctx }
    }
  });
  return {
    c() {
      create_component(onscrollappear.$$.fragment);
    },
    m(target, anchor) {
      mount_component(onscrollappear, target, anchor);
      current = true;
    },
    p(ctx2, dirty) {
      const onscrollappear_changes = {};
      if (dirty & /*$$scope, formations*/
      1025) {
        onscrollappear_changes.$$scope = { dirty, ctx: ctx2 };
      }
      onscrollappear.$set(onscrollappear_changes);
    },
    i(local) {
      if (current)
        return;
      transition_in(onscrollappear.$$.fragment, local);
      current = true;
    },
    o(local) {
      transition_out(onscrollappear.$$.fragment, local);
      current = false;
    },
    d(detaching) {
      destroy_component(onscrollappear, detaching);
    }
  };
}
function create_default_slot_1(ctx) {
  let h2;
  return {
    c() {
      h2 = element("h2");
      h2.textContent = "Exp\xE9riences professionnelles";
      attr(h2, "class", "svelte-1vov8fz");
    },
    m(target, anchor) {
      insert(target, h2, anchor);
    },
    p: noop,
    d(detaching) {
      if (detaching) {
        detach(h2);
      }
    }
  };
}
function create_else_block(ctx) {
  let p;
  return {
    c() {
      p = element("p");
      p.textContent = "Chargement des exp\xE9riences...";
    },
    m(target, anchor) {
      insert(target, p, anchor);
    },
    p: noop,
    i: noop,
    o: noop,
    d(detaching) {
      if (detaching) {
        detach(p);
      }
    }
  };
}
function create_if_block_1(ctx) {
  let each_1_anchor;
  let current;
  let each_value = ensure_array_like(
    /*experiences*/
    ctx[2]
  );
  let each_blocks = [];
  for (let i = 0; i < each_value.length; i += 1) {
    each_blocks[i] = create_each_block2(get_each_context2(ctx, each_value, i));
  }
  const out = (i) => transition_out(each_blocks[i], 1, 1, () => {
    each_blocks[i] = null;
  });
  return {
    c() {
      for (let i = 0; i < each_blocks.length; i += 1) {
        each_blocks[i].c();
      }
      each_1_anchor = empty();
    },
    m(target, anchor) {
      for (let i = 0; i < each_blocks.length; i += 1) {
        if (each_blocks[i]) {
          each_blocks[i].m(target, anchor);
        }
      }
      insert(target, each_1_anchor, anchor);
      current = true;
    },
    p(ctx2, dirty) {
      if (dirty & /*experiences*/
      4) {
        each_value = ensure_array_like(
          /*experiences*/
          ctx2[2]
        );
        let i;
        for (i = 0; i < each_value.length; i += 1) {
          const child_ctx = get_each_context2(ctx2, each_value, i);
          if (each_blocks[i]) {
            each_blocks[i].p(child_ctx, dirty);
            transition_in(each_blocks[i], 1);
          } else {
            each_blocks[i] = create_each_block2(child_ctx);
            each_blocks[i].c();
            transition_in(each_blocks[i], 1);
            each_blocks[i].m(each_1_anchor.parentNode, each_1_anchor);
          }
        }
        group_outros();
        for (i = each_value.length; i < each_blocks.length; i += 1) {
          out(i);
        }
        check_outros();
      }
    },
    i(local) {
      if (current)
        return;
      for (let i = 0; i < each_value.length; i += 1) {
        transition_in(each_blocks[i]);
      }
      current = true;
    },
    o(local) {
      each_blocks = each_blocks.filter(Boolean);
      for (let i = 0; i < each_blocks.length; i += 1) {
        transition_out(each_blocks[i]);
      }
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(each_1_anchor);
      }
      destroy_each(each_blocks, detaching);
    }
  };
}
function create_if_block(ctx) {
  let p;
  let t;
  return {
    c() {
      p = element("p");
      t = text(
        /*error*/
        ctx[1]
      );
      attr(p, "class", "error svelte-1vov8fz");
    },
    m(target, anchor) {
      insert(target, p, anchor);
      append(p, t);
    },
    p(ctx2, dirty) {
      if (dirty & /*error*/
      2)
        set_data(
          t,
          /*error*/
          ctx2[1]
        );
    },
    i: noop,
    o: noop,
    d(detaching) {
      if (detaching) {
        detach(p);
      }
    }
  };
}
function create_default_slot2(ctx) {
  let experiencecard;
  let t;
  let current;
  experiencecard = new experienceCard_default({
    props: { experience: (
      /*experience*/
      ctx[5]
    ) }
  });
  return {
    c() {
      create_component(experiencecard.$$.fragment);
      t = space();
    },
    m(target, anchor) {
      mount_component(experiencecard, target, anchor);
      insert(target, t, anchor);
      current = true;
    },
    p(ctx2, dirty) {
      const experiencecard_changes = {};
      if (dirty & /*experiences*/
      4)
        experiencecard_changes.experience = /*experience*/
        ctx2[5];
      experiencecard.$set(experiencecard_changes);
    },
    i(local) {
      if (current)
        return;
      transition_in(experiencecard.$$.fragment, local);
      current = true;
    },
    o(local) {
      transition_out(experiencecard.$$.fragment, local);
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(t);
      }
      destroy_component(experiencecard, detaching);
    }
  };
}
function create_each_block2(ctx) {
  let onscrollappear;
  let current;
  onscrollappear = new OnScrollAppear_default({
    props: {
      index: (
        /*index*/
        ctx[7]
      ),
      $$slots: { default: [create_default_slot2] },
      $$scope: { ctx }
    }
  });
  return {
    c() {
      create_component(onscrollappear.$$.fragment);
    },
    m(target, anchor) {
      mount_component(onscrollappear, target, anchor);
      current = true;
    },
    p(ctx2, dirty) {
      const onscrollappear_changes = {};
      if (dirty & /*$$scope, experiences*/
      1028) {
        onscrollappear_changes.$$scope = { dirty, ctx: ctx2 };
      }
      onscrollappear.$set(onscrollappear_changes);
    },
    i(local) {
      if (current)
        return;
      transition_in(onscrollappear.$$.fragment, local);
      current = true;
    },
    o(local) {
      transition_out(onscrollappear.$$.fragment, local);
      current = false;
    },
    d(detaching) {
      destroy_component(onscrollappear, detaching);
    }
  };
}
function create_fragment5(ctx) {
  let section0;
  let onscrollappear0;
  let t0;
  let current_block_type_index;
  let if_block0;
  let t1;
  let section1;
  let onscrollappear1;
  let t2;
  let current_block_type_index_1;
  let if_block1;
  let current;
  onscrollappear0 = new OnScrollAppear_default({
    props: {
      $$slots: { default: [create_default_slot_3] },
      $$scope: { ctx }
    }
  });
  const if_block_creators = [create_if_block_2, create_if_block_3, create_else_block_1];
  const if_blocks = [];
  function select_block_type(ctx2, dirty) {
    if (
      /*error*/
      ctx2[1]
    )
      return 0;
    if (
      /*formations*/
      ctx2[0].length > 0
    )
      return 1;
    return 2;
  }
  current_block_type_index = select_block_type(ctx, -1);
  if_block0 = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);
  onscrollappear1 = new OnScrollAppear_default({
    props: {
      $$slots: { default: [create_default_slot_1] },
      $$scope: { ctx }
    }
  });
  const if_block_creators_1 = [create_if_block, create_if_block_1, create_else_block];
  const if_blocks_1 = [];
  function select_block_type_1(ctx2, dirty) {
    if (
      /*error*/
      ctx2[1]
    )
      return 0;
    if (
      /*experiences*/
      ctx2[2].length > 0
    )
      return 1;
    return 2;
  }
  current_block_type_index_1 = select_block_type_1(ctx, -1);
  if_block1 = if_blocks_1[current_block_type_index_1] = if_block_creators_1[current_block_type_index_1](ctx);
  return {
    c() {
      section0 = element("section");
      create_component(onscrollappear0.$$.fragment);
      t0 = space();
      if_block0.c();
      t1 = space();
      section1 = element("section");
      create_component(onscrollappear1.$$.fragment);
      t2 = space();
      if_block1.c();
      attr(section0, "class", "svelte-1vov8fz");
      attr(section1, "class", "svelte-1vov8fz");
    },
    m(target, anchor) {
      insert(target, section0, anchor);
      mount_component(onscrollappear0, section0, null);
      append(section0, t0);
      if_blocks[current_block_type_index].m(section0, null);
      insert(target, t1, anchor);
      insert(target, section1, anchor);
      mount_component(onscrollappear1, section1, null);
      append(section1, t2);
      if_blocks_1[current_block_type_index_1].m(section1, null);
      current = true;
    },
    p(ctx2, [dirty]) {
      const onscrollappear0_changes = {};
      if (dirty & /*$$scope*/
      1024) {
        onscrollappear0_changes.$$scope = { dirty, ctx: ctx2 };
      }
      onscrollappear0.$set(onscrollappear0_changes);
      let previous_block_index = current_block_type_index;
      current_block_type_index = select_block_type(ctx2, dirty);
      if (current_block_type_index === previous_block_index) {
        if_blocks[current_block_type_index].p(ctx2, dirty);
      } else {
        group_outros();
        transition_out(if_blocks[previous_block_index], 1, 1, () => {
          if_blocks[previous_block_index] = null;
        });
        check_outros();
        if_block0 = if_blocks[current_block_type_index];
        if (!if_block0) {
          if_block0 = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx2);
          if_block0.c();
        } else {
          if_block0.p(ctx2, dirty);
        }
        transition_in(if_block0, 1);
        if_block0.m(section0, null);
      }
      const onscrollappear1_changes = {};
      if (dirty & /*$$scope*/
      1024) {
        onscrollappear1_changes.$$scope = { dirty, ctx: ctx2 };
      }
      onscrollappear1.$set(onscrollappear1_changes);
      let previous_block_index_1 = current_block_type_index_1;
      current_block_type_index_1 = select_block_type_1(ctx2, dirty);
      if (current_block_type_index_1 === previous_block_index_1) {
        if_blocks_1[current_block_type_index_1].p(ctx2, dirty);
      } else {
        group_outros();
        transition_out(if_blocks_1[previous_block_index_1], 1, 1, () => {
          if_blocks_1[previous_block_index_1] = null;
        });
        check_outros();
        if_block1 = if_blocks_1[current_block_type_index_1];
        if (!if_block1) {
          if_block1 = if_blocks_1[current_block_type_index_1] = if_block_creators_1[current_block_type_index_1](ctx2);
          if_block1.c();
        } else {
          if_block1.p(ctx2, dirty);
        }
        transition_in(if_block1, 1);
        if_block1.m(section1, null);
      }
    },
    i(local) {
      if (current)
        return;
      transition_in(onscrollappear0.$$.fragment, local);
      transition_in(if_block0);
      transition_in(onscrollappear1.$$.fragment, local);
      transition_in(if_block1);
      current = true;
    },
    o(local) {
      transition_out(onscrollappear0.$$.fragment, local);
      transition_out(if_block0);
      transition_out(onscrollappear1.$$.fragment, local);
      transition_out(if_block1);
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(section0);
        detach(t1);
        detach(section1);
      }
      destroy_component(onscrollappear0);
      if_blocks[current_block_type_index].d();
      destroy_component(onscrollappear1);
      if_blocks_1[current_block_type_index_1].d();
    }
  };
}
function instance5($$self, $$props, $$invalidate) {
  let formations = [];
  let error = "";
  async function getFormations() {
    try {
      const response = await fetch("/formations");
      if (!response.ok) {
        throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
      }
      const data = await response.json();
      $$invalidate(0, formations = data);
    } catch (e) {
      console.error("\xC9chec de la r\xE9cup\xE9ration des formations :", e.message || e);
      $$invalidate(1, error = "Impossible de charger les formations. Veuillez r\xE9essayer plus tard.");
    }
  }
  let experiences = [];
  async function getExperiences() {
    try {
      const response = await fetch("/experiences");
      if (!response.ok) {
        throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
      }
      const data = await response.json();
      $$invalidate(2, experiences = data);
    } catch (e) {
      console.error("\xC9chec de la r\xE9cup\xE9ration des exp\xE9riences :", e.message || e);
      $$invalidate(1, error = "Impossible de charger les exp\xE9riences. Veuillez r\xE9essayer plus tard.");
    }
  }
  onMount(() => {
    getFormations();
    getExperiences();
  });
  return [formations, error, experiences];
}
var Cv = class extends SvelteComponent {
  constructor(options) {
    super();
    init(this, options, instance5, create_fragment5, safe_not_equal, {}, add_css3);
  }
};
create_custom_element(Cv, {}, [], [], true);
var cv_default = Cv;

// ihm/App.svelte
function add_css4(target) {
  append_styles(target, "svelte-18wtlvy", `@import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css");@import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");.banner.svelte-18wtlvy.svelte-18wtlvy{position:relative;height:calc(100vh - var(--header-height));display:flex;flex-direction:column;overflow:hidden;background:url('../assets/background.webp') center/cover no-repeat}.banner.svelte-18wtlvy.svelte-18wtlvy::before{content:'';position:absolute;top:0;left:0;right:0;bottom:0;background-image:inherit;background-size:cover;background-position:center;z-index:0}.banner-content.svelte-18wtlvy.svelte-18wtlvy{position:relative;z-index:1;color:white;text-align:left;text-shadow:2px 2px 4px rgba(0, 0, 0, 0.9);min-width:0;max-width:100%;padding-left:20px}.banner-content.svelte-18wtlvy h1.svelte-18wtlvy,.banner-content.svelte-18wtlvy h2.svelte-18wtlvy{margin:0;padding:0.2em 0;text-align:left;transition:opacity 0.5s ease-in-out}.banner-content.svelte-18wtlvy h2.fade-out.svelte-18wtlvy{opacity:0}.parcours.svelte-18wtlvy.svelte-18wtlvy{align-self:center;margin-top:auto;margin-bottom:0.5em;z-index:1;color:white;font-size:1.5rem;text-shadow:2px 2px 4px rgba(0, 0, 0, 0.9)}.scroll-indicator.svelte-18wtlvy.svelte-18wtlvy{align-self:center;margin-bottom:1rem;font-size:1.5rem;color:white;text-shadow:2px 2px 4px rgba(0, 0, 0, 0.9);z-index:1;animation:svelte-18wtlvy-bounce 2s infinite}@keyframes svelte-18wtlvy-bounce{0%,20%,50%,80%,100%{transform:translateY(0)}40%{transform:translateY(-10px)}60%{transform:translateY(-5px)}}.its-me.svelte-18wtlvy.svelte-18wtlvy{display:flex;z-index:1;margin:calc(20/100 * (100vh - var(--header-height)) + var(--header-height)) auto auto 0}.my-pic.svelte-18wtlvy.svelte-18wtlvy{margin-left:2rem;width:clamp(80px, 15vw, 200px);height:clamp(80px, 15vw, 200px);background-color:whitesmoke;border-radius:50%;border:3px solid rgba(255, 255, 255, 0.7);box-shadow:0 0 0 5px rgba(255, 255, 255, 0.3),
		            0 4px 10px rgba(0, 0, 0, 0.5);object-fit:cover;transition:all 0.3s ease}.my-pic.svelte-18wtlvy.svelte-18wtlvy:hover{transform:scale(1.05);box-shadow:0 0 0 8px rgba(255, 255, 255, 0.4),
		            0 6px 15px rgba(0, 0, 0, 0.6)}.metier.svelte-18wtlvy.svelte-18wtlvy,.name.svelte-18wtlvy.svelte-18wtlvy{white-space:normal;word-wrap:break-word}`);
}
function create_fragment6(ctx) {
  let div3;
  let div1;
  let img;
  let img_src_value;
  let t0;
  let div0;
  let h1;
  let t2;
  let h2;
  let t3;
  let t4;
  let span;
  let t6;
  let div2;
  let t7;
  let cv;
  let current;
  cv = new cv_default({});
  return {
    c() {
      div3 = element("div");
      div1 = element("div");
      img = element("img");
      t0 = space();
      div0 = element("div");
      h1 = element("h1");
      h1.textContent = `${name}`;
      t2 = space();
      h2 = element("h2");
      t3 = text(
        /*currentH2Text*/
        ctx[0]
      );
      t4 = space();
      span = element("span");
      span.textContent = "Mon parcours";
      t6 = space();
      div2 = element("div");
      div2.innerHTML = `<i class="bi bi-arrow-down-circle"></i>`;
      t7 = space();
      create_component(cv.$$.fragment);
      attr(img, "class", "my-pic svelte-18wtlvy");
      if (!src_url_equal(img.src, img_src_value = "./assets/icons8-male-user.svg"))
        attr(img, "src", img_src_value);
      attr(img, "alt", "");
      attr(h1, "class", "name svelte-18wtlvy");
      attr(h2, "class", "metier svelte-18wtlvy");
      toggle_class(
        h2,
        "fade-out",
        /*isFading*/
        ctx[1]
      );
      attr(div0, "class", "banner-content svelte-18wtlvy");
      attr(div1, "class", "its-me svelte-18wtlvy");
      attr(span, "class", "parcours svelte-18wtlvy");
      attr(div2, "class", "scroll-indicator svelte-18wtlvy");
      attr(div3, "class", "banner svelte-18wtlvy");
    },
    m(target, anchor) {
      insert(target, div3, anchor);
      append(div3, div1);
      append(div1, img);
      append(div1, t0);
      append(div1, div0);
      append(div0, h1);
      append(div0, t2);
      append(div0, h2);
      append(h2, t3);
      append(div3, t4);
      append(div3, span);
      append(div3, t6);
      append(div3, div2);
      insert(target, t7, anchor);
      mount_component(cv, target, anchor);
      current = true;
    },
    p(ctx2, [dirty]) {
      if (!current || dirty & /*currentH2Text*/
      1)
        set_data(
          t3,
          /*currentH2Text*/
          ctx2[0]
        );
      if (!current || dirty & /*isFading*/
      2) {
        toggle_class(
          h2,
          "fade-out",
          /*isFading*/
          ctx2[1]
        );
      }
    },
    i(local) {
      if (current)
        return;
      transition_in(cv.$$.fragment, local);
      current = true;
    },
    o(local) {
      transition_out(cv.$$.fragment, local);
      current = false;
    },
    d(detaching) {
      if (detaching) {
        detach(div3);
        detach(t7);
      }
      destroy_component(cv, detaching);
    }
  };
}
var name = "Cl\xE9ment Calia";
var inge = "Expert en ing\xE9nierie logicielle";
var fullstack = "D\xE9veloppeur Fullstack";
var transitionDuration = 500;
var displayDuration = 3e3;
function instance6($$self, $$props, $$invalidate) {
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
    init(this, options, instance6, create_fragment6, safe_not_equal, {}, add_css4);
  }
};
customElements.define("index-portfolio", create_custom_element(App, {}, [], [], true));
var App_default = App;
export {
  App_default as default
};
