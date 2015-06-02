#!/usr/bin/env bash

pkgbuild --root sources \
    --identifier org.pfc.aramirez.agent \
    --version 1.0.6 \
    --ownership recommended \
    --scripts scripts \
    pkg/output.pkg && productbuild --distribution distribution.xml \
    --resources resources \
    --package-path pkg \
    --version 1.0.5 \
    pkg/pfcaramirezagent-1.0.6.pkg

