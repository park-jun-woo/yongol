package stml

// sitemapFullGrammarSrc is the Phase001 plan's full sitemap grammar example,
// the fixture for TestParseSitemapReader_FullGrammar.

const sitemapFullGrammarSrc = `
<nav data-sitemap data-layout="app">
  <ul>
    <li data-page="dashboard" data-index>대시보드</li>
    <li>건물 관리
      <ul>
        <li data-page="building-list">건물 목록
          <ul>
            <li data-page="building-detail">건물 상세</li>
          </ul>
        </li>
      </ul>
    </li>
    <li data-page="member-list" data-icon="users" data-menu="false">멤버</li>
    <li><a href="https://docs.example.com">사용자 매뉴얼</a></li>
  </ul>
</nav>
<nav data-sitemap data-layout="bare" data-entry>
  <ul>
    <li data-page="login">로그인</li>
  </ul>
</nav>`
