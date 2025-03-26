<template>
  <div id="app">
    <nav class="main-nav">
      <router-link to="/" class="nav-logo">XbookLab</router-link>
      <div class="nav-links">
        <router-link to="/">首页</router-link>
        <router-link to="/profile">个人中心</router-link>
      </div>
    </nav>
    <router-view/>
  </div>
</template>

<style scoped>
#app {
  font-family: Arial, sans-serif;
}
.main-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 40px;
  background: #f8f9fa;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
.nav-logo {
  font-size: 24px;
  font-weight: bold;
  color: #2c3e50;
  text-decoration: none;
}
.nav-links a {
  margin-left: 30px;
  color: #2c3e50;
  text-decoration: none;
}
.nav-links a.router-link-exact-active {
  color: #42b983;
  font-weight: bold;
}
</style>