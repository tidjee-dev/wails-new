package embedded

const ExternalLinkErrorTS = `export function handleExternalLinkError(err: unknown) {
  if (typeof err === "string" && err === "external_url_blocked") {
    alert("This link is not allowed.");
    return;
  }

  console.error("Failed to open external link:", err);
  alert("Unable to open link.");
}
`

const UseExternalLinkTS = `import { OpenExternalURL } from '../../wailsjs/go/main/App';
import { handleExternalLinkError } from "./externalLinkError";

export function useExternalLink() {
  return (url: string) => {
    return async (e: MouseEvent | KeyboardEvent) => {
      e.preventDefault();

      try {
        await OpenExternalURL(url);
      } catch (err) {
        handleExternalLinkError(err);
      }
    };
  };
}
`
