<template>
  <header>
    <nav class="container">
      <div class="branding">
        <router-link class="header" :to="{ name: 'Home' }"
          >Data-Blog</router-link
        >
        <div class="nav-links">
          <ul>
            <router-link class="link" to="#">Home</router-link>
            <router-link class="link" to="#">Blog</router-link>
            <router-link class="link" to="#">Create Post</router-link>
            <router-link class="link" to="#">Login/Register</router-link>
          </ul>
        </div>
      </div>
    </nav>
    <menuIcon />
    <transition name="mobile-nav">
      <ul>
        <router-link class="link" to="#">Home</router-link>
        <router-link class="link" to="#">Blog</router-link>
        <router-link class="link" to="#">Create Post</router-link>
        <router-link class="link" to="#">Login/Register</router-link>
      </ul>
    </transition>
  </header>
</template>

<script>
import { menuIcon } from "../assets/Icons/bars-regular.svg"

export default {
  name: "navigation",
  components: {
    menuIcon,
  },
}
</script>

<style lang="scss" scoped>
header {
  background-color: #fff;
  padding: 0 25px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1),
    0 2px 4px -1px rgba(0, 0, 0, 0.06);
  z-index: 99;

  .link {
    font-weight: 500;
    padding: 0 8px;
    transition: 0.3s color ease;

    &:hover {
      color: #1eb8b8;
    }
  }

  nav {
    display: flex;
    padding: 25px 0;

    .branding {
      display: flex;
      align-items: center;

      .header {
        font-weight: 600;
        font-size: 24px;
        color: #000;
        text-decoration: none;
      }
    }
  }
}
</style>
