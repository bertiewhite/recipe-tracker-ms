<script>
  let message = "loading"

  fetch('http://localhost:8080/hello')
          .then(response => response.json())
          .then(data => message = data.message);

</script>

<main>
  <p>
    {message}
  </p>
</main>

<style>

</style>