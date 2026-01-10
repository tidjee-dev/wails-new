import Heading from "@theme/Heading";
import clsx from "clsx";
import { Cpu, Layers, Rocket } from "lucide-react";
import type { ReactNode } from "react";
import styles from "./styles.module.css";

type FeatureItem = {
	id: string;
	title: string;
	Icon: React.ComponentType<{ className?: string }>;
	description: ReactNode;
};

const FeatureList: FeatureItem[] = [
	{
		id: "easy",
		title: "Easy to Use",
		Icon: Rocket,
		description: (
			<>Instantly scaffold a modern Wails app with zero boilerplate.</>
		),
	},
	{
		id: "structure",
		title: "Focus on What Matters",
		Icon: Layers,
		description: (
			<>Clean frontend and backend structure, ready for real projects.</>
		),
	},
	{
		id: "wails",
		title: "Powered by Wails",
		Icon: Cpu,
		description: <>Native desktop apps with Go and a modern web stack.</>,
	},
];

function Feature({ title, Icon, description }: FeatureItem) {
	return (
		<article className={clsx("col col--4", styles.featureCol)}>
			<div className={styles.featureCard}>
				<div className={styles.iconWrapper}>
					<Icon className={styles.featureIcon} aria-hidden="true" />
				</div>

				<Heading as="h3" className={styles.featureTitle}>
					{title}
				</Heading>

				<p className={styles.featureDescription}>{description}</p>
			</div>
		</article>
	);
}

export default function HomepageFeatures(): ReactNode {
	return (
		<section className={styles.features} aria-labelledby="features-title">
			<div className="container">
				<Heading as="h2" id="features-title" className={styles.sectionTitle}>
					Why wails-new?
				</Heading>

				<div className="row">
					{FeatureList.map((feature) => (
						<Feature key={feature.id} {...feature} />
					))}
				</div>
			</div>
		</section>
	);
}
