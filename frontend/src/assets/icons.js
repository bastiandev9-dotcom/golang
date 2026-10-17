const DEVICON = (path) =>
  `https://cdn.jsdelivr.net/gh/devicons/devicon/icons/${path}`

const overrides = {
  html5: DEVICON('html5/html5-original.svg'),
  css3: DEVICON('css3/css3-original.svg'),
  adobephotoshop: DEVICON('photoshop/photoshop-original.svg'),
  adobeillustrator: DEVICON('illustrator/illustrator-original.svg'),
  linkedin: DEVICON('linkedin/linkedin-original.svg'),
  twitter: DEVICON('twitter/twitter-original.svg'),
}

export const resolveIcon = (key, color = 'FFFFFF') =>
  overrides[key] || `https://cdn.simpleicons.org/${key}/${color}`
