<template>
  <v-app>
    <v-app-bar absolute app color="white">
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="./assets/logo.jpg"
          transition="scale-transition"
          width="105"
        />
      </div>

      <v-toolbar-title>UniqueSSO</v-toolbar-title>
    </v-app-bar>

    <Particles id="tsparticles" :options="particlesOpts" />

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",
  data: () => {
    return {
      particlesOpts: {
        background: {},
        fpsLimit: 60,
        interactivity: {
          detectsOn: "window",
          events: {
            onClick: {
              enable: false,
            },
            onHover: {
              enable: false,
            },
            resize: true,
          },
        },
        particles: {
          color: {
            value: "#4B6587",
            opacity: 0.3,
          },
          links: {
            color: "#C8C6C6",
            distance: 100,
            enable: true,
            opacity: 0.2,
            width: 3,
          },
          collisions: {
            enable: true,
          },
          move: {
            direction: "none",
            enable: true,
            outMode: "bounce",
            random: false,
            speed: 3,
            straight: false,
          },
          number: {
            density: {
              enable: true,
              value_area: 700,
            },
            value: 75,
          },
          opacity: {
            value: 0.5,
          },
          shape: {
            type: "circle",
          },
          size: {
            random: true,
            value: 5,
          },
        },
        detectRetina: true,
      },
    };
  },
};
</script>

<style>
#tsparticles {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>