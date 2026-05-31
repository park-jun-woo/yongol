package stml

const richPageHTML = `<main data-route="/things/:id">
  <article data-fetch="GetThing" data-param-id="route.id">
    <div class="wrapper">
      <h2 data-field="Name"></h2>
      <div data-component="Avatar" data-field="Owner" class="av"></div>
      <span data-bind="Status"></span>
      <p class="plain">just static text</p>
    </div>
    <section data-each="Items">
      <div class="row">
        <div class="cell">
          <span data-bind="Title"></span>
          <em class="muted">label</em>
        </div>
      </div>
    </section>
    <div data-component="DirectCard" data-field="Card" class="dc"></div>
    <form data-action="UpdateThing">
      <fieldset>
        <input data-field="Title" type="text" placeholder="title" class="inp" />
        <div data-component="Picker" data-field="Tag" class="pk">
          <input data-field="Inner" type="text" />
        </div>
        <div class="static-action-wrap">
          <input data-field="Nested" type="text" />
          <div data-component="WalkCard" data-field="Walk" class="wc">
            <div data-component="InnerWalk" data-field="Inner2" class="iw">
              <input data-field="DeepNested" type="text" />
            </div>
          </div>
          <button type="submit">Save</button>
        </div>
      </fieldset>
    </form>
  </article>
  <footer class="page-footer">
    <p>static footer</p>
  </footer>
</main>`

const layoutHTML = `<div>
  <nav>
    <a data-nav="/home">Home</a>
  </nav>
  <slot data-outlet />
</div>`
