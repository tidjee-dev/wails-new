export function handleExternalLinkError(err: unknown) {
  if (typeof err === "string" && err === "external_url_blocked") {
    alert("This link is not allowed.");
    return;
  }

  console.error("Failed to open external link:", err);
  alert("Unable to open link.");
}
