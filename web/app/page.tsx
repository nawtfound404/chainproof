"use client"

import { useState, useEffect } from "react"
import styles from "./page.module.css"

type MerkleProofItem = {
  hash: string
  position: string
}

type ProofResponse = {
  hash: string
  cid: string
  tx_hash: string
  merkle_root: string
  merkle_proof: MerkleProofItem[]
  timestamp: string
  encrypted: boolean
}

// Minimal Animated Node Graph Background Component
function NetworkGraphBackground() {
  return (
    <div style={{ position: 'absolute', top: 0, left: 0, width: '100%', height: '100%', overflow: 'hidden', zIndex: -1, opacity: 0.15 }}>
      <svg width="100%" height="100%" xmlns="http://www.w3.org/2000/svg">
        <defs>
          <radialGradient id="nodeGlow" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stopColor="var(--accent-cyan)" stopOpacity="1" />
            <stop offset="100%" stopColor="transparent" stopOpacity="0" />
          </radialGradient>
        </defs>
        <line x1="10%" y1="20%" x2="40%" y2="80%" stroke="var(--text-secondary)" strokeWidth="1" opacity="0.5" />
        <line x1="40%" y1="80%" x2="80%" y2="40%" stroke="var(--text-secondary)" strokeWidth="1" opacity="0.5" />
        <line x1="80%" y1="40%" x2="90%" y2="90%" stroke="var(--text-secondary)" strokeWidth="1" opacity="0.5" />
        <line x1="10%" y1="20%" x2="80%" y2="40%" stroke="var(--text-secondary)" strokeWidth="1" opacity="0.2" />

        <circle cx="10%" cy="20%" r="4" fill="var(--text-primary)" />
        <circle cx="40%" cy="80%" r="6" fill="var(--accent-cyan)" filter="drop-shadow(0 0 8px var(--accent-cyan))" />
        <circle cx="80%" cy="40%" r="5" fill="var(--accent-emerald)" filter="drop-shadow(0 0 8px var(--accent-emerald))" />
        <circle cx="90%" cy="90%" r="3" fill="var(--text-primary)" />
      </svg>
    </div>
  )
}

// Educational Tooltip Component
const Tooltip = ({ text, children }: { text: string, children: React.ReactNode }) => {
  const [show, setShow] = useState(false)
  return (
    <span
      className={styles.tooltipContainer}
      onMouseEnter={() => setShow(true)}
      onMouseLeave={() => setShow(false)}
    >
      <span className={styles.tooltipTarget}>{children}</span>
      {show && (
        <div className={styles.tooltip}>
          {text}
        </div>
      )}
    </span>
  )
}

function MerkleTreeVisualizer({
  proof,
  leaf,
  root
}: {
  proof: MerkleProofItem[],
  leaf: string,
  root: string
}) {
  if (!proof || proof.length === 0) return null

  // Reconstruct the path visually from bottom (leaf) to top (root)
  // Level 0: Leaf + Sibling
  // Level N: Root

  return (
    <div style={{ marginTop: '2rem' }}>
      <h3 style={{ fontSize: '1rem', fontWeight: 600, marginBottom: '1rem', color: 'var(--text-primary)' }}>
        Cryptographic Proof Path Selection
      </h3>
      <div className={styles.merkleContainer}>
        {/* We render from top (root) to bottom (leaves) visually by using flex column-reverse in CSS */}

        <div className={styles.merkleLevel}>
          {/* Root Level */}
          <div className={`${styles.merkleNode} ${styles.isProofPath}`} title="Merkle Root">
            {root.substring(0, 10)}...{root.substring(root.length - 8)}
          </div>
        </div>

        {proof.map((p, index) => {
          // For each proof item, it's a sibling to our current computed hash path
          // This is simplified for visual effect; a true full tree render needs all leaves
          return (
            <div key={index} className={styles.merkleLevel}>
              {p.position === 'left' ? (
                <>
                  <div className={`${styles.merkleNode}`} title="Sibling Hash">
                    {p.hash.substring(0, 10)}...
                  </div>
                  <div className={`${styles.merkleNode} ${styles.isProofPath}`} title="Path Hash">
                    (Computed)
                  </div>
                </>
              ) : (
                <>
                  <div className={`${styles.merkleNode} ${styles.isProofPath}`} title="Path Hash">
                    (Computed)
                  </div>
                  <div className={`${styles.merkleNode}`} title="Sibling Hash">
                    {p.hash.substring(0, 10)}...
                  </div>
                </>
              )}
            </div>
          )
        })}

        <div className={styles.merkleLevel}>
          {/* Target Leaf */}
          <div className={`${styles.merkleNode} ${styles.isTarget}`} title="Target Leaf Hash">
            {leaf.substring(0, 10)}...{leaf.substring(leaf.length - 8)}
          </div>
          {/* Target's sibling is handled in the proof loop level 0 usually, but visually we just place leaf at bottom */}
        </div>

      </div>
    </div>
  )
}

const PIPELINE_STEPS = [
  { id: 1, title: "Canonicalizing Data", desc: "Sorting keys and standardizing JSON structure format", tooltip: "Canonicalization ensures that logically equivalent data always produces the exact same byte sequence before hashing." },
  { id: 2, title: "Computing SHA256 Hash", desc: "Generating deterministic cryptographic leaf hash", tooltip: "A one-way cryptographic hash function that produces a unique 256-bit signature for your canonicalized data." },
  { id: 3, title: "Uploading to IPFS", desc: "Pinning canonical format to decentralized storage", tooltip: "The InterPlanetary File System (IPFS) ensures your data is permanently available and content-addressable." },
  { id: 4, title: "Adding to Merkle Batch", desc: "Inserting leaf hash into real-time Merkle tree", tooltip: "Merkle Trees allow us to batch thousands of proofs into a single root hash, saving immense gas costs on Ethereum." },
  { id: 5, title: "Anchoring Root On Ethereum", desc: "Publishing batch Merkle root to Sepolia smart contract", tooltip: "The Merkle Root is published to an Ethereum Smart Contract, permanently anchoring the State of the batch to the blockchain." },
  { id: 6, title: "Generating Inclusion Proof", desc: "Constructing cryptographic path verification proof", tooltip: "An inclusion proof is a sequence of sibling hashes that proves your specific transaction was included in the anchored Merkle Root." },
]

function PipelineVisualizer({ currentStep }: { currentStep: number }) {
  if (currentStep === 0) return null

  return (
    <section className={styles.card}>
      <div className={styles.cardHeader}>
        <h2 className={styles.cardTitle}>Live Cryptographic Pipeline</h2>
        <p className={styles.cardDescription}>Real-time infrastructure processing visualization</p>
      </div>

      <div className={styles.pipeline}>
        {PIPELINE_STEPS.map((step) => {
          const isActive = currentStep === step.id
          const isCompleted = currentStep > step.id

          return (
            <div
              key={step.id}
              className={`${styles.pipelineStep} ${isActive ? styles.active : ''} ${isCompleted ? styles.completed : ''}`}
            >
              <div className={styles.stepLine}></div>
              <div className={styles.stepIcon}>
                {isCompleted ? (
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                ) : isActive ? (
                  <div className={styles.spinner}></div>
                ) : (
                  <span style={{ fontWeight: 600 }}>{step.id}</span>
                )}
              </div>
              <div className={styles.stepContent}>
                <div className={styles.stepTitle}>
                  <Tooltip text={step.tooltip}>
                    {step.title}
                  </Tooltip>
                </div>
                <div className={styles.stepDescription}>
                  {step.desc}
                </div>
              </div>
            </div>
          )
        })}
      </div>
    </section>
  )
}

export default function Home() {
  const [input, setInput] = useState('{\n  "event": "transfer",\n  "amount": 1000,\n  "recipient": "0x123..."\n}')
  const [result, setResult] = useState<ProofResponse | null>(null)
  const [pipelineStep, setPipelineStep] = useState(0) // 0 means not started
  const [error, setError] = useState<string | null>(null)
  const [verifyStatus, setVerifyStatus] = useState<string | null>(null)

  const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:9000"

  const handleSubmit = async () => {
    setError(null)
    setResult(null)
    setVerifyStatus(null)
    setPipelineStep(1) // Start pipeline

    try {
      const parsed = JSON.parse(input)

      // Simulate pipeline UI progress
      for (let i = 1; i <= 6; i++) {
        setPipelineStep(i)
        // Wait realistic but fast simulation time
        await new Promise(r => setTimeout(r, i === 3 ? 800 : i === 5 ? 1200 : 600))
      }

      // Actually fetch the data (in background during simulation or after)
      const res = await fetch(`${BACKEND_URL}/proof`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ data: parsed }),
      })

      if (!res.ok) {
        throw new Error("Proof generation failed on backend")
      }

      const data = await res.json()
      setResult(data)
      setPipelineStep(7) // Complete

    } catch (err: any) {
      setError(err.message || "Invalid JSON or server error")
      setPipelineStep(0)
    }
  }

  const handleVerify = async () => {
    if (!result) return

    try {
      setVerifyStatus("Verifying...")

      const url = `${BACKEND_URL}/verify?hash=${result.hash.replace(/^0x/, "")}`

      const res = await fetch(url)

      if (!res.ok) {
        throw new Error(`Verification failed: ${res.status}`)
      }

      const data = await res.json()

      if (data.on_chain === true) {
        setVerifyStatus("✔ Proof Anchored On-Chain")
      } else {
        setVerifyStatus("❌ Not Found On-Chain")
      }
    } catch (err: any) {
      setVerifyStatus(`❌ Verification Error: ${err.message || err}`)
    }
  }

  return (
    <main>
      <div className={styles.container}>

        {/* 1. Hero Section */}
        <section className={styles.hero}>
          <NetworkGraphBackground />
          <h1 className={styles.heroTitle}>ChainProof</h1>
          <div className={styles.heroSubtitle}>Deterministic Web3 Proof & Audit Infrastructure</div>
          <p className={styles.heroDescription}>
            Transform off-chain data into verifiable, tamper-proof cryptographic commitments anchored on Ethereum.
          </p>
        </section>

        {/* 2. Off-Chain Data Input Section */}
        <section className={styles.card}>
          <div className={styles.cardHeader}>
            <h2 className={styles.cardTitle}>
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"></path><polyline points="3.27 6.96 12 12.01 20.73 6.96"></polyline><line x1="12" y1="22.08" x2="12" y2="12"></line></svg>
              Simulate Off-Chain Event
            </h2>
            <p className={styles.cardDescription}>
              Your data will be canonicalized, hashed, batched into a Merkle tree, and anchored on Ethereum.
            </p>
          </div>

          <textarea
            className={styles.textarea}
            value={input}
            onChange={(e) => setInput(e.target.value)}
            spellCheck={false}
            disabled={pipelineStep > 0 && pipelineStep < 7}
          />

          <div style={{ marginTop: '1.5rem', display: 'flex', justifyContent: 'flex-end', alignItems: 'center', gap: '1rem' }}>
            <div style={{ fontSize: '0.875rem', color: 'var(--text-secondary)', display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
              <span style={{ display: 'inline-block', width: 8, height: 8, borderRadius: '50%', background: 'var(--accent-cyan)', boxShadow: 'var(--accent-glow)' }}></span>
              Ready to canonicalize
            </div>
            <button
              onClick={handleSubmit}
              disabled={pipelineStep > 0 && pipelineStep < 7}
              className={`${styles.btn} ${styles.btnPrimary}`}
            >
              {(pipelineStep > 0 && pipelineStep < 7) ? "Initializing..." : "Generate Commitment"}
            </button>
          </div>
        </section>

        {/* 3. Pipeline Visualizer */}
        <PipelineVisualizer currentStep={pipelineStep} />

        {/* 4. Proof Result Card */}
        {error && (
          <div className={styles.card} style={{ borderColor: 'rgba(239, 68, 68, 0.3)', background: 'rgba(239, 68, 68, 0.05)' }}>
            <div style={{ color: '#ef4444', fontWeight: 500 }}>{error}</div>
          </div>
        )}

        {result && pipelineStep === 7 && (
          <section className={styles.card} style={{ borderColor: 'rgba(16, 185, 129, 0.3)' }} id="proof-result">
            <div className={styles.cardHeader}>
              <h2 className={styles.cardTitle}>
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" style={{ color: 'var(--accent-emerald)' }}><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path><polyline points="22 4 12 14.01 9 11.01"></polyline></svg>
                Commitment Generated Successfully
              </h2>
            </div>

            <div style={{ display: 'grid', gap: '1rem', background: 'rgba(0,0,0,0.2)', padding: '1.5rem', borderRadius: '8px', border: '1px solid var(--surface-border)' }}>
              <div>
                <div style={{ fontSize: '0.875rem', color: 'var(--text-secondary)', marginBottom: '0.25rem' }}>Leaf Hash</div>
                <div style={{ fontFamily: 'var(--font-mono)', color: 'var(--accent-cyan)', fontSize: '0.875rem', wordBreak: 'break-all' }}>{result.hash}</div>
              </div>
              <div>
                <div style={{ fontSize: '0.875rem', color: 'var(--text-secondary)', marginBottom: '0.25rem' }}>IPFS CID</div>
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.875rem' }}>
                  <a href={`https://ipfs.io/ipfs/${result.cid}`} target="_blank" rel="noreferrer" style={{ color: 'var(--text-primary)', textDecoration: 'underline' }}>{result.cid}</a>
                </div>
              </div>
              <div>
                <div style={{ fontSize: '0.875rem', color: 'var(--text-secondary)', marginBottom: '0.25rem' }}>Merkle Root</div>
                <div style={{ fontFamily: 'var(--font-mono)', color: 'var(--accent-emerald)', fontSize: '0.875rem', wordBreak: 'break-all' }}>{result.merkle_root}</div>
              </div>
              <div>
                <div style={{ fontSize: '0.875rem', color: 'var(--text-secondary)', marginBottom: '0.25rem' }}>Ethereum Transaction</div>
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.875rem' }}>
                  <a href={`https://sepolia.etherscan.io/tx/${result.tx_hash}`} target="_blank" rel="noreferrer" style={{ color: 'var(--text-primary)', textDecoration: 'underline' }}>{result.tx_hash}</a>
                </div>
              </div>
            </div>

            {/* Merkle Tree Visualizer */}
            {result.merkle_proof && result.merkle_proof.length > 0 && (
              <MerkleTreeVisualizer
                proof={result.merkle_proof}
                leaf={result.hash}
                root={result.merkle_root}
              />
            )}

            <div style={{ marginTop: '1.5rem', display: 'flex', alignItems: 'center', gap: '1rem' }}>
              <button onClick={handleVerify} className={`${styles.btn} ${styles.btnSuccess}`}>
                Verify On-Chain
              </button>
              {verifyStatus && (
                <div style={{ fontWeight: 500, color: verifyStatus.startsWith('✔') ? 'var(--accent-emerald)' : verifyStatus.startsWith('❌') ? '#ef4444' : 'var(--text-secondary)' }}>
                  {verifyStatus}
                </div>
              )}
            </div>
          </section>
        )}

      </div>
    </main>
  )
}