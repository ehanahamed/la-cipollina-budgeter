<script>
    import { base } from "$app/paths";
    import {Calendar, DayGrid} from '@event-calendar/core';
    import BackArrowIcon from "$lib/icons/BackArrow.svelte";
    import CheckmarkIcon from "$lib/icons/Checkmark.svelte";
    import { dateToYMDString } from "$lib/dateToYMDString.js";
    let { data } = $props();
    let events = [];
    console.log(data)
    data.weeks.forEach((week) => {
        if (week.firstDate && week.lastDate) {
            const [startY, startM, startD] = week.startDate.split("-");
            const startDate = new Date(startY, startM - 1, startD);
            /* "start" is the first monday, "first" is the actual first day,
            if they're closed on monday, then first might be a different day,
            but start will still be monday */
            const [firstY, firstM, firstD] = week.firstDate.split("-");
            const firstDate = new Date(firstY, firstM - 1, firstD);
            const [lastY, lastM, lastD] = week.lastDate.split("-");
            const lastDate = new Date(lastY, lastM - 1, lastD);
            events.push({
                title: `Week of ${[
                    "Jan", "Feb", "Mar", "Apr",
                    "May", "Jun", "Jul", "Aug",
                    "Sep", "Oct", "Nov", "Dec"
                ][startDate.getMonth()]} ${startDate.getDate()}`,
                start: firstDate,
                end: lastDate
            });
        }
    })

    let options = $state({
        view: 'dayGridMonth',
        events: events,
        headerToolbar: {
            start: "title",
            center: "",
            end: "prev,next"
        },
        firstDay: localStorage.getItem("budgeter:calendar.firstWeekDay") == "mon" ?
            1 : 0,
        eventClick: (info) => {
            console.log(info.event)
        },
        eventBackgroundColor: "var(--main)"
    });
</script>
<style>
:global {
    .ec-title {
        font-weight: normal;
    }
    .ec-button:focus {
        background-color: var(--bg-3);
    }
    .ec-event-time {
        display: none;
    }
    .ec-event-title {
        font-size: 1rem;
    }
    .ec-event-body {
        cursor: pointer;
    }
}
</style>
<div class="grid page" style="margin-top: 2rem; margin-bottom: 10rem;">
    <div class="content" style="padding-top: 2rem;">
    <div>
        <a href={base} class="button faint">
            <BackArrowIcon></BackArrowIcon>
            Back
        </a>
    </div>
<p class="center">Select a week report to view</p>
<div class="no-child-margin-top no-child-clickable-effect">
<Calendar plugins={[DayGrid]} {options} />
</div>
<div style="margin-top: 4rem;">
    <h4 style="font-weight: normal;">Calendar Settings</h4>
    <p>First day of the week:</p>
    <div class="combo-select">
        <button class="left {
            localStorage.getItem("budgeter:calendar.firstWeekDay") == "mon" ?
                "" : "selected"
        }" onclick={() => {
            localStorage.setItem("budgeter:calendar.firstWeekDay", "sun");
            window.location.reload();
        }}>
            <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
            Sunday
        </button>
        <button class="right {
            localStorage.getItem("budgeter:calendar.firstWeekDay") == "mon" ?
                "selected" : ""
        }" onclick={() => {
            localStorage.setItem("budgeter:calendar.firstWeekDay", "mon");
            window.location.reload();
        }}>
            <CheckmarkIcon class="combo-selected-icon"></CheckmarkIcon>
            Monday
        </button>
    </div>
</div>
</div>
</div>

