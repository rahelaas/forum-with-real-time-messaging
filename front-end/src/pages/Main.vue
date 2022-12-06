<template>

  <map name="image_map">
    <area
        alt="LogIn"
        title="Log In"
        href="/login"
        coords="1177,888,1245,955"
        shape="rect"
    />
    <area
        alt="Sign Up"
        title="Sign Up"
        href="/signup"
        coords="318,885,393,956"
        shape="rect"
    />
  </map>
  <h1>Welcome to the Forum!</h1>
</template>

<script>
export default {
  // eslint-disable-next-line
  name: "Main",
};
</script>

<style scoped>
.matrix {
  max-width: 100%;
  max-height: 100%;
  bottom: 0;
  left: 0;
  margin: auto;
  overflow: auto;
  position: absolute;
  right: 0;
  top: 0;
  -o-object-fit: contain;
  object-fit: contain;
}
</style>
