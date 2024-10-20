<svelte:head>
  <title>Voima - Workout Planner</title>
  <meta name="description" content="The Voima app login page." />
</svelte:head>

<script>
    let count = 0;
  
    const increment = () => {
      count += 1;
    };
  </script>
  
  <h1>Voima</h1>
  
<p>Rep count: {count}</p>
<button on:click={increment}>Increment</button>