<script>
  let name = '';
  let mood = null;
  let showGreeting = false;

  $: timeOfDay = (() => {
    const hour = new Date().getHours();
    if (hour < 12) return 'morning';
    if (hour < 18) return 'afternoon';
    return 'evening';
  })();

  $: greetingMessage = (() => {
    const messages = {
      happy: `Good ${timeOfDay}, ${name}! Your happiness is contagious today! ✨`,
      tired: `Good ${timeOfDay}, ${name}. Take it easy, rest is important. 💤`,
      excited: `Good ${timeOfDay}, ${name}! Your excitement is electrifying! ⚡`,
      grumpy: `Good ${timeOfDay}, ${name}. Hope your day gets better soon. ☁️`,
      peaceful: `Good ${timeOfDay}, ${name}. May your calm carry you through. 🍃`
    };
    return messages[mood] || '';
  })();

  $: moodGradient = (() => {
    const gradients = {
      happy: 'linear-gradient(135deg, #FCD34D 0%, #F59E0B 100%)',
      tired: 'linear-gradient(135deg, #818CF8 0%, #6366F1 100%)',
      excited: 'linear-gradient(135deg, #F472B6 0%, #EC4899 100%)',
      grumpy: 'linear-gradient(135deg, #94A3B8 0%, #64748B 100%)',
      peaceful: 'linear-gradient(135deg, #86EFAC 0%, #34D399 100%)'
    };
    return gradients[mood] || 'linear-gradient(135deg, #F3F4F6 0%, #E5E7EB 100%)';
  })();

   const moods = [
    { id: 'happy', emoji: '😊', label: 'Happy' },
    { id: 'tired', emoji: '😴', label: 'Tired' },
    { id: 'excited', emoji: '🤩', label: 'Excited' },
    { id: 'grumpy', emoji: '😠', label: 'Grumpy' },
    { id: 'peaceful', emoji: '😌', label: 'Peaceful' }
  ];

  // Event handlers - just regular functions
  function handleGreet() {
    if (name && mood) {
      showGreeting = true;  // Just assign! No setState
    }
  }

  function handleReset() {
    showGreeting = false;  // Just assign! No setState
  }
</script>

<div class="container" style="background: {moodGradient}">
  
  {#if !showGreeting}
    <!-- FORM VIEW -->
    <div class="card">
      <h1>How are you feeling today?</h1>

      <!-- Name Input -->
      <div class="input-group">
        <label for="name">Your Name</label>
        <input
          id="name"
          type="text"
          bind:value={name}
          placeholder="Enter your name..."
        />
      </div>

      <!-- Mood Selector -->
      <div class="mood-group">
        <label>Your Mood</label>
        <div class="mood-buttons">
          {#each moods as m}
            <button
              class="mood-btn"
              class:selected={mood === m.id}
              on:click={() => mood = m.id}
            >
              <span class="emoji">{m.emoji}</span>
              <span class="label">{m.label}</span>
            </button>
          {/each}
        </div>
      </div>

      <!-- Greet Button -->
      <button
        class="greet-btn"
        disabled={!name || !mood}
        on:click={handleGreet}
      >
        Greet Me! 🎉
      </button>
    </div>

  {:else}
    <!-- GREETING VIEW -->
    <div class="greeting">
      <div class="emoji-large">
        {moods.find(m => m.id === mood)?.emoji}
      </div>
      <h1 class="message">{greetingMessage}</h1>
      <button class="reset-btn" on:click={handleReset}>
        Change Mood
      </button>
    </div>
  {/if}
  
</div>