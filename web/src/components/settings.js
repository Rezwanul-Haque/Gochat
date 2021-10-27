import { createClient, createMicrophoneAndCameraTracks } from "agora-rtc-react";

const appId = "f253930f26f549b1997f141cd305f9a8";

const token = window.localStorage.getItem("rtc_token") ? window.localStorage.getItem("rtc_token") : null;

console.log("settings token", token)

export const config = { mode: "rtc", codec: "vp8", appId: appId, token: token };
export const useClient = createClient(config);
export const useMicrophoneAndCameraTracks = createMicrophoneAndCameraTracks();
