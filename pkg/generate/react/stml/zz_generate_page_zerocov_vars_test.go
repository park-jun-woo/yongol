package stml

const richPageHTML = `<main class="container">
  <section data-fetch="ListReservations" data-param-status="route.Status">
    <ul data-each="reservations">
      <li>
        <span>{item.title}</span>
      </li>
    </ul>
    <button data-action="DeleteReservation" data-param-id="route.ID">삭제</button>
  </section>
  <article data-fetch="GetReservation" data-param-id="route.ID">
    <footer data-state="canCancel">
      <form data-action="UpdateReservation">
        <input data-field="title" />
        <input data-field="note" />
        <button>저장</button>
      </form>
      <button data-action="CancelReservation">취소</button>
    </footer>
  </article>
  <div data-static="info">
    <p>static content</p>
    <button data-action="Login">로그인</button>
  </div>
</main>`
