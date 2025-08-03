<script>
    import { onMount } from "svelte";
    import { base } from "$app/paths";
    import { goto } from "$app/navigation";
    import "vanillajs-datepicker/css/datepicker.css";
    import "$lib/vanillajs-datepicker-ehui.css";
    import BackIcon from "$lib/icons/BackArrow.svelte";
    import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
    import { dateToYMDString } from "$lib/dateToYMDString.js";
    let { data } = $props();
    let datePickerInput;
    let datePicker;
    onMount(async () => {
        const { Datepicker } = await import("vanillajs-datepicker");
        datePicker = new Datepicker(
            datePickerInput
        );
        datePicker.setDate(new Date());
    });
    let selectedDateYMD;
</script>
<div class="grid page" style="margin-top: 4rem; margin-bottom: 8rem;">
    <div class="content">
        <div class="flex">
            <a href="{base}" class="button faint">
                <BackIcon></BackIcon>
                Back
            </a>
        </div>
        <p>Select Day to View/Edit:</p>
        <div class="eh-datepicker-container" bind:this={datePickerInput}></div>
        <div class="flex">
            <button onclick={async () => {
                selectedDateYMD = dateToYMDString(
                    datePicker.getDate()
                );
                try {
                    const res = await fetch(
                        `${data.PUBLIC_API_URL}/days/${
                            selectedDateYMD
                        }`, {
                        method: "GET",
                        headers: {
                            "Authorization": `Bearer ${
                                localStorage.getItem("budgeter:auth")
                            }`
                        }
                    });
                    const resJson = await res.json();
                    if (resJson.id) {
                        goto(`${base}/days/${selectedDateYMD}`);
                    } else if (res.status == 404) {
                        goto(`${base}/record-day?d=${selectedDateYMD}`);
                    } else {
                        console.log(resJson);
                        alert("hmm something didn't work")
                    }
                } catch (err) {
                    console.error(err);
                    alert("something went wrong?");
                }
            }}>
                <CheckmarkIcon></CheckmarkIcon>
                Select
            </button>
        </div>
    </div>
</div>
