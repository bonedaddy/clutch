import React from "react";
import styled, { keyframes } from "styled-components";

const rotate = keyframes`
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(-360deg);
  }
`;

const StyledSvg = styled.svg`
  height: 40px;
  width: 40px;
  &:hover {
    animation: ${rotate} 5s linear;
  }
`;

const Logo: React.FC = () => (
  <StyledSvg id="logo" viewBox="0 0 250 250">
    <defs>
      <clipPath id="clutch_svg__b" clipPathUnits="userSpaceOnUse">
        <path d="M0 1500h1500V0H0z" />
      </clipPath>
      <clipPath id="clutch_svg__a" clipPathUnits="userSpaceOnUse">
        <rect
          transform="scale(1 -1)"
          ry={0.075}
          y={-675.814}
          x={494.261}
          height={187.5}
          width={187.5}
          opacity={0.8}
          fill="#e6e6e6"
          strokeWidth={0.75}
        />
      </clipPath>
    </defs>
    <path d="M-659.015 901.085h2000v-2000h-2000z" fill="none" />
    <g clipPath="url(#clutch_svg__a)" transform="matrix(1.33333 0 0 -1.33333 -659.015 901.085)">
      <g clipPath="url(#clutch_svg__b)">
        <path
          d="M853.19 617.03a4.185 4.185 0 015.269-2.698l30.426 9.82a4.185 4.185 0 11-2.571 7.964l-30.425-9.818a4.184 4.184 0 01-2.698-5.268M1329.569 314.774l-22.966-22.966-9.865 9.865 22.966 22.967a6.934 6.934 0 004.932 2.042 6.938 6.938 0 004.933-2.042c2.719-2.72 2.719-7.146 0-9.866m-133.454-119.922v14.692c0 8.102 6.59 14.693 14.692 14.693s14.693-6.59 14.693-14.692v-14.692zm62.228-26.04h-95.072v17.67h95.072zm32.476 126.943l9.865-9.865-68.274-68.275a23.183 23.183 0 01-8.696 11.033zm-99.006 12.465h10.276a4.185 4.185 0 010 8.37h-10.276v14.274l6.287 6.286h25.416l6.286-6.286v-33.38h-37.99zm143.674 22.338a15.242 15.242 0 01-10.851 4.495 15.244 15.244 0 01-10.851-4.495l-98.406-98.407a23.115 23.115 0 01-4.572.456c-12.716 0-23.062-10.345-23.062-23.063v-14.692h-14.533v87.286a6.983 6.983 0 006.975 6.977h61.24a6.94 6.94 0 005.91-3.29 4.186 4.186 0 017.096 4.44c-2.83 4.521-7.692 7.22-13.005 7.22h-3.255v35.113c0 1.11-.442 2.174-1.226 2.96l-8.74 8.737a4.181 4.181 0 01-2.958 1.225h-28.883c-1.11 0-2.175-.44-2.96-1.225l-8.738-8.737a4.192 4.192 0 01-1.226-2.96v-35.114h-3.255c-8.461 0-15.346-6.884-15.346-15.346v-87.286h-5.755a4.185 4.185 0 01-4.185-4.185v-26.042a4.185 4.185 0 014.185-4.185h103.443a4.185 4.185 0 014.185 4.185v26.042a4.185 4.185 0 01-4.185 4.185h-5.755v10.718a4.185 4.185 0 01-8.37 0v-10.718h-14.533v12.387l101.616 101.617c5.984 5.982 5.984 15.72 0 21.702M1324.636 209.734h-26.875a4.186 4.186 0 110-8.37h26.875a4.185 4.185 0 110 8.37M1324.636 189.272h-26.875a4.186 4.186 0 110-8.37h26.875a4.185 4.185 0 110 8.37M1324.636 168.811h-26.875a4.186 4.186 0 110-8.37h26.875a4.185 4.185 0 110 8.37M626.767 507.54l-.51 29.283a59.816 59.816 0 0113.36 15.998l13.992-23.29a84.727 84.727 0 00-26.842-21.992m-42.555-9.354l14.192 25.6a59.012 59.012 0 0119.583 7.201l.473-27.209a83.736 83.736 0 00-34.248-5.592m-70.879 45.371l29.285.51a59.812 59.812 0 0115.998-13.359l-23.29-13.993a84.732 84.732 0 00-21.993 26.842m36.017 113.432l.51-29.284a59.8 59.8 0 01-13.359-15.998l-13.993 23.292a84.735 84.735 0 0026.842 21.99m42.555 9.354l-14.192-25.599a58.968 58.968 0 01-19.583-7.202l-.473 27.21a83.738 83.738 0 0034.248 5.591m70.88-45.37l-29.286-.509a59.83 59.83 0 01-15.998 13.358l23.292 13.993a84.73 84.73 0 0021.991-26.842m-16.246-28.363a58.945 58.945 0 01-7.203 19.584l27.21.473a83.746 83.746 0 005.592-34.248zm-13.128 60.535l-25.057-15.054a59.127 59.127 0 01-20.294 3.574c-.09 0-.179-.007-.269-.007l13.186 23.786a83.641 83.641 0 0032.434-12.299m-116.23-25.529l15.052-25.057a59.122 59.122 0 01-3.573-20.294c0-.09.006-.179.006-.269l-23.786 13.186a83.643 83.643 0 0012.3 32.434m-13.29-45.35c0 1.288.033 2.57.091 3.845l25.6-14.192a58.983 58.983 0 017.202-19.583l-27.21-.473a83.737 83.737 0 00-5.683 30.402m38.818-70.88l25.056 15.054a59.115 59.115 0 0120.296-3.574c.09 0 .178.006.268.007l-13.186-23.786a83.641 83.641 0 00-32.434 12.3m45.352 19.85c-28.138 0-51.03 22.891-51.03 51.03 0 28.137 22.892 51.028 51.03 51.028 28.137 0 51.03-22.89 51.03-51.029 0-28.138-22.893-51.03-51.03-51.03m70.88 5.678l-15.054 25.057a59.128 59.128 0 013.574 20.295c0 .09-.007.18-.007.268l23.786-13.186a83.657 83.657 0 00-12.299-32.434m-70.88 137.892c-51.027 0-92.54-41.513-92.54-92.54 0-51.027 41.513-92.54 92.54-92.54 51.026 0 92.54 41.513 92.54 92.54 0 51.027-41.514 92.54-92.54 92.54"
          fill="#02acbe"
        />
        <path
          d="M588.059 572.38c-5.45 0-9.884 4.433-9.884 9.884 0 5.45 4.434 9.884 9.884 9.884s9.884-4.434 9.884-9.884-4.434-9.885-9.884-9.885m0 28.14c-10.066 0-18.255-8.189-18.255-18.255 0-10.067 8.189-18.255 18.255-18.255 10.066 0 18.255 8.188 18.255 18.255 0 10.066-8.19 18.255-18.255 18.255M609.664 614.985a4.184 4.184 0 01-5.919 0 4.183 4.183 0 010-5.918l11.116-11.116a4.172 4.172 0 012.96-1.226 4.186 4.186 0 012.96 7.145zM561.256 566.577a4.184 4.184 0 01-5.919 0 4.183 4.183 0 010-5.918l11.116-11.116a4.172 4.172 0 012.96-1.226 4.186 4.186 0 012.96 7.145zM614.861 566.577l-11.116-11.115a4.186 4.186 0 015.919-5.92l11.116 11.117a4.185 4.185 0 01-5.919 5.918M558.297 596.725c1.07 0 2.142.41 2.96 1.226l11.115 11.115a4.184 4.184 0 11-5.92 5.92l-11.114-11.117a4.183 4.183 0 012.959-7.144M570.074 292.41l1.758 14.392c.44 3.588 3.86 8.166 7.177 9.604l46.445 20.127c.934.404 2.087.618 3.336.618 2.76 0 5.57-1.027 7.16-2.617l33.658-33.658c2.39-2.391 3.343-7.393 2-10.496l-20.128-46.444c-1.438-3.317-6.016-6.74-9.606-7.177l-14.389-1.76a9.55 9.55 0 00-1.144-.064c-3.86 0-8.784 1.898-11.207 4.32L574.33 280.06c-2.689 2.689-4.717 8.576-4.256 12.35m-38.637-121.79a6.144 6.144 0 00-4.37-1.81 6.142 6.142 0 00-4.37 1.81l-17.002 17.002a6.188 6.188 0 000 8.74l70.247 70.248 25.742-25.742zm77.778 62.718c3.988-3.988 11.03-6.774 17.125-6.774h.001c.743 0 1.47.043 2.16.127l14.39 1.759c6.514.796 13.66 6.137 16.27 12.158l20.126 46.444c2.71 6.254 1.058 14.925-3.76 19.742l-33.659 33.658c-3.173 3.175-8.062 5.069-13.078 5.069-2.384 0-4.688-.452-6.664-1.308l-46.444-20.126c-6.02-2.61-11.362-9.756-12.158-16.27l-1.759-14.39c-.772-6.316 2.147-14.787 6.647-19.286l1.612-1.612-70.247-70.248c-5.674-5.673-5.674-14.905 0-20.577l17-17.001a14.46 14.46 0 0110.29-4.263c3.886 0 7.54 1.515 10.288 4.263l70.248 70.247z"
          fill="#02acbe"
        />
      </g>
    </g>
  </StyledSvg>
);

export default Logo;