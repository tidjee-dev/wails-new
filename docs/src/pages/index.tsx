import Link from "@docusaurus/Link";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import HomepageFeatures from "@site/src/components/HomepageFeatures";
import Heading from "@theme/Heading";
import Layout from "@theme/Layout";
import clsx from "clsx";
import type { ReactNode } from "react";
import styles from "./index.module.css";

function HomepageHeader() {
	const { siteConfig } = useDocusaurusContext();

	return (
		<header className={clsx("hero hero--primary", styles.heroBanner)}>
			<div className={clsx("container", styles.heroContainer)}>
				{/* Left */}
				<div className={styles.heroContent}>
					<Heading as="h1" className={styles.heroTitle}>
						{siteConfig.title}
					</Heading>
					<p className={styles.heroSubtitle}>{siteConfig.tagline}</p>

					<div className={styles.getStartedButton}>
						<Link
							className="button button--link button--secondary button--lg"
							to="/docs/getting-started/installation"
						>
							&gt; Get Started &lt;
						</Link>
					</div>
				</div>

				{/* Right */}
				<div className={styles.heroMedia}>
					<img
						src={require("@site/static/img/wails-new_demo.webp").default}
						alt="Animated demo of the wails-new CLI scaffolding a project"
						className={styles.heroImage}
						loading="eager"
					/>
				</div>
			</div>
		</header>
	);
}

export default function Home(): ReactNode {
	return (
		<Layout
			title="Home"
			description="Bootstrap a modern Wails app with Svelte 5, Vite, and Tailwind CSS"
		>
			<HomepageHeader />
			<main>
				<HomepageFeatures />
			</main>
		</Layout>
	);
}
