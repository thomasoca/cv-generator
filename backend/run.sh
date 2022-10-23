#!/bin/bash

# Declare an array of string with type
declare -a StringArray=("pdfx" "biber" "bibhang" "biblabelsep" "pubtype" "bibsetup" "bibitemsep" "trimclip")
 
# Iterate the string array using for loop
for val in ${StringArray[@]}; do
   sed -i "/\b$val\b/d" altacv.cls
done

# Add modification to altacv.cls
cat <<EOT >> altacv.cls
% Modification
\newcommand{\cvproject}[3]{%
  {\large\color{emphasis}#1\par}
  \smallskip\normalsize
  \ifstrequal{#2}{}{}{
  \textbf{\color{accent}#2}\par
  \smallskip}
  \ifstrequal{#3}{}{}{{\small\makebox[0.5\linewidth][l]{\faCalendar~#3}}}%
  \medskip\normalsize
}
\newcommand{\cvskillstr}[2]{%
  \textcolor{emphasis}{\textbf{#1}}\hfill
  \textbf{\color{body}#2}\par
}
EOT
