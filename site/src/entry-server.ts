import "abortcontroller-polyfill/dist/abortcontroller-polyfill-only";

import { renderToString } from "@vue/server-renderer";
import { makeApp } from "~/main";
import { renderSSRHead } from "@unhead/ssr";

export async function render(url: string) {
  // page is being rendered at this point
  const { app, router, head } = await makeApp({});
  await router.push(url);

  const ctx: any = {};
  const body = await renderToString(app, ctx);
  const ssrHead: SSRHeadPayload = await renderSSRHead(head);

  return { head: ssrHead, body: body };
}

function ssrRender(
  url: string,
  moduleEntryPoint: string,
  cssEntryPoint: string,
) {
  return render(url).then(({ head, body }) => {
    let html = `
<!DOCTYPE html<!--htmlAttrs-->>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="icon" href="/assets/favicon.svg" type="image/svg+xml">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Dosis:wght@200..800&display=swap" rel="stylesheet">
  <!--moduleEntryPoint-->
  <!--cssEntryPoint-->
  <!--headTags-->
  <!--preload-links-->
</head>
<body <!--bodyAttrs-->>
  <div id="app"><!--app-html--></div>
  <!--bodyTags-->
  <script>
    (function() {
      const prefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
      const setting = localStorage.getItem('color-schema') || 'auto'
      if (setting === 'dark' || (prefersDark && setting !== 'light'))
        document.documentElement.classList.toggle('dark', true)
    })()
  </script>
</body>
</html>
`;
    Object.entries(head).forEach(([key, value]) => {
      html = html.replace(`<!--${key}-->`, value);
    });
    html = html.replace(`<!--app-html-->`, body);
    html = html.replace(
      `<!--moduleEntryPoint-->`,
      `<script type="module" crossorigin src="${moduleEntryPoint}"></script>`,
    );
    html = html.replace(
      `<!--cssEntryPoint-->`,
      `<link rel="stylesheet" href="${cssEntryPoint}">`,
    );
    return html;
  });
}

(globalThis as any).ssrRender = ssrRender;

export interface SSRHeadPayload {
  headTags: string;
  bodyTags: string;
  bodyTagsOpen: string;
  htmlAttrs: string;
  bodyAttrs: string;
}
