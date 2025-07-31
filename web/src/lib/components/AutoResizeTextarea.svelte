<script>
    let { value = $bindable(), div, textarea } = $props();
    let textareaElement;

    function resize() {
        textareaElement.style.height = "auto";
        textareaElement.style.height = textareaElement.scrollHeight + "px";
    }

    $effect(() => resize());
</script>

<div {...div}>
    <textarea
        bind:value={value}
        bind:this={textareaElement}
        on:input={resize}
        {...textarea}
        style="overflow:hidden;resize:none;{textarea?.style ?? ""}"
    ></textarea>
</div>
