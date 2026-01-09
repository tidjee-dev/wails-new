import { OpenExternalURL } from '../../wailsjs/go/main/App';
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
