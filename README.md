# Vanilla OS - מנהל חבילות Apx

Apx הוא מנהל החבילות ברירת המחדל של Vanilla OS. הוא עטיפה סביב מנהלי חבילות מרובים להתקנת חבילות והפעלת פקודות בתוך מיכל מנוהל.

## סקירה כללית

**Apx** (/à·peks/) מספק ממשק אחיד לניהול חבילות על פני הפצות Linux שונות. הוא משתמש ב-[distrobox](https://github.com/89luca89/distrobox) כדי ליצור סביבות בטחונות שבהן ניתן להתקין חבילות ללא השפעה על מערכת המארח שלך.

## תכונות

- **תמיכה בהפצות מרובות**: התקנת חבילות מהפצות Linux שונות
- **מבוסס על קונטיינר**: חבילות מותקנות בקונטיינרים מבודדים
- **הפשטת מנהל החבילות**: ממשק אחיד עבור מנהלי חבילות שונים (apt, dnf, pacman, וכו')
- **שילוב שולחן עבודה**: ייצוא אוטומטי של רשומות שולחן עבודה לאפליקציות מותקנות
- **ניהול תת-מערכות**: יצירה וניהול מרובות תת-מערכות עם מחסניות שונות

## הערימה

- **שפה**: Go 1.25.0
- **Framework**: Vanilla OS SDK + CLI builder
- **ממשק משתמש**: Charmbracelet (bubbletea, lipgloss) לממשק טרמינל
- **קונטיינר**: שילוב distrobox

## מבנה הפרויקט

```
cmd/
  main.go              # נקודת כניסה של היישום
  main_check.go        # וריאנט אימות מחרוזות
  locales/             # תמיכה בשפות מרובות (33+ שפות)

internal/cli/
  structs.go           # הגדרות פקודות CLI
  runtime.go           # מטפלי ביצוע תת-מערכות
  subsystems.go        # ניהול מחזור חיי תת-מערכות
  stacks.go            # פעולות ערימה
  pkgmanagers.go       # פעולות מנהל חבילות

core/                  # לוגיקה עסקית ליבית (תת-מערכות, ערימות, ניהול pkg)
config/                # קבצי תצורה
distrobox/             # סקריפטי שילוב distrobox
```

## בנייה

### דרישות מוקדמות

- Go 1.25.0 או גבוה יותר
- make
- podman או docker
- git

### שלבי בנייה

```bash
# שיבוט וכניסה לתיקייה
git clone --recursive https://github.com/yourusername/Vanilla-.git
cd Vanilla-

# בנייה
make build

# התקנה לרמת המערכת
sudo make install
sudo make install-manpages

# התקנה למיקום מותאם אישית
make install PREFIX=$HOME/.local
make install-manpages PREFIX=$HOME/.local
```

## שימוש

### פקודות בסיסיות

```bash
# רישום תת-מערכות זמינות
apx subsystems list

# יצירת תת-מערכת חדשה
apx subsystems new --name=ubuntu-latest

# כניסה לתת-מערכת
apx ubuntu-latest enter

# התקנת חבילות בתת-מערכת
apx ubuntu-latest install package-name

# הפעלת פקודה בתת-מערכת
apx ubuntu-latest run command

# רישום ערימות זמינות
apx stacks list

# ניהול מנהלי חבילות
apx pkgmanagers list
```

## גישה מרחוק

Apx תומך בגישה מרחוק דרך SSH, מה שמאפשר ניהול חבילות וביצוע פקודות בשרתים וציוד מרוחק.

### הגדרת SSH

```bash
# חיבור לשרת מרוחק
ssh user@remote-host

# הרצת פקודות Apx מרחוק
ssh user@remote-host "apx subsystems list"

# התקנת חבילות דרך SSH
ssh user@remote-host "apx ubuntu-latest install package-name"
```

### שימוש ב-iPhone או כלים ניידים

עבור גישה מרחוק מ-iPhone או מכשירים ניידים אחרים, תוכל להשתמש בתוכנות SSH כמו:

- **Termius** - ממשק SSH מלא עם תמיכה בקונטיינרים
- **iSH** - קונסולה SSH עבור iOS
- **Prompt** - קליינט SSH מקצועי לiOS

```bash
# דוגמה: חיבור מ-iPhone דרך Termius
ssh user@your-server.com
apx ubuntu-latest enter
```

### גישה מרחוק מאובטחת

```bash
# הגדרת מפתח SSH ציבורי לגישה מאובטחת
ssh-copy-id -i ~/.ssh/id_rsa.pub user@remote-host

# כניסה ללא סיסמה
ssh user@remote-host "apx subsystems list"

# הפעלת סקריפט מרחוק
ssh user@remote-host < /path/to/local/script.sh
```

## תמיכה ב-Android

Apx תומך בניהול חבילות וביצוע פקודות דרך מכשירי Android בעזרת Termux וSSH.

### Termux Integration

Termux היא טרמינל מלא עבור Android שתומכת בלימוד Apx בצורה ישירה.

```bash
# התקנת Termux מ-F-Droid
# https://f-droid.org/en/packages/com.termux/

# התקנת Apx ב-Termux
pkg install go git
git clone https://github.com/yourusername/Vanilla-.git
cd Vanilla-
make build
sudo make install
```

### גישה ל-Apx דרך Android

```bash
# חיבור לשרת Apx משרת האחסון שלך
ssh user@your-server.com

# הרצת פקודות Apx מרחוק
ssh user@your-server.com "apx subsystems list"

# התקנת חבילות מ-Android
ssh user@your-server.com "apx ubuntu-latest install python3"
```

### יישומי Android SSH

שימוש ביישומי SSH פופולריים עבור Android:

- **Termux** - שורת פקודה ישירה וSSH
- **Termius** - קליינט SSH מקצועי
- **JuiceSSH** - ממשק SSH משופר
- **ConnectBot** - קליינט SSH קל משקל

```bash
# דוגמה: חיבור דרך Termius
1. פתח Termius
2. צור חיבור SSH חדש
3. הזן את פרטי השרת שלך
4. הפעל: apx subsystems list
```

### פיתוח אפליקציות Android עם Apx

```bash
# התקנת כלי פיתוח Android בתת-מערכת
apx ubuntu-latest install android-sdk android-ndk

# קומפילציה של אפליקציות דרך SSH מ-Android
ssh user@dev-server.com "apx android-dev compile --project /path/to/project"

# בדיקה של עבודה עם ADB מרחוק
ssh user@dev-server.com "apx android-dev adb shell"
```

### סנכרון קבצים דרך Android

```bash
# העברת קבצים מ-Android לשרת Apx
scp /sdcard/myproject/* user@server.com:/home/user/projects/

# משיכת תוצאות מ-Apx ל-Android
scp user@server.com:/build/output/* /sdcard/Downloads/
```

## AI ויישומי בינה מלאכותית

Apx משתלב יכולות בינה מלאכותית כדי לשפר את חווית ניהול החבילות וההתקנה.

### המלצות חבילות בחכמה

Apx משתמש בלימוד מכונה כדי להמליץ על חבילות רלוונטיות על סמך:
- רישום התקנה קודם שלך
- תבניות שימוש נפוצות
- תלויויות ופרויקטים דומים

```bash
# קבל המלצות AI עבור חבילה
apx ai-recommend package-name

# ניתוח חבילה עם AI
apx ai-analyze ubuntu-latest
```

### פתרון בעיות בעזרת AI

AI יכול לסייע בטיפול בשגיאות וביצירת סביבה מאופטימלת:

```bash
# קבל אבחון AI לבעיות התקנה
apx ai-diagnose

# קבל הצעות AI לפתרון בעיות
apx ai-troubleshoot error-message

# קבל ייעוץ AI להגדרה אופטימלית
apx ai-optimize
```

### הסכמה חכמה של תלויויות

AI מנתח תלויויות בחבילות ומצא:
- תלויויות חסרות
- ניגודים בגרסאות
- חבילות מיותרות

```bash
# בדוק תלויויות עם AI
apx ai-check-deps subsystem-name

# קבל הצעות להסרת חבילות מיותרות
apx ai-cleanup-deps subsystem-name

# קבל דוח AI על בריאות המערכת
apx ai-health-report
```

### רצועת זמן נתמכת

AI תומך בעברית, אנגלית ועוד 30+ שפות:

```bash
# שנה שפה לעברית
apx --language=he ai-recommend package-name

# קבל סיוע AI בשפתך
apx ai-help --lang=en
```

## תלויויות

### תלויויות ישירות
- `github.com/google/uuid` - יצירת UUID
- `github.com/vanilla-os/sdk` - Framework יישום
- `gopkg.in/yaml.v2` - תצורת YAML

### תלויויות ממשק משתמש
- `github.com/charmbracelet/bubbletea` - Framework ממשק טרמינל
- `github.com/charmbracelet/lipgloss` - עיצוב טרמינל

### תלויויות AI
- `github.com/tmc/langchain-go` - שיתו�� פעולה LLM
- `github.com/sashabaranov/go-openai` - שיתוף פעולה OpenAI (אופציונלי)
- `github.com/google/generative-ai-python` - שיתוף פעולה Google AI (אופציונלי)

## תיעוד

לתיעוד מפורט, בקר ב: https://docs.vanillaos.org/docs/en/apx

## תרגומים

תרום תרגומים דרך [Weblate](https://hosted.weblate.org/projects/vanilla-os/apx)

## רישיון

GNU General Public License v3.0

## מחברים

- Mirko Brombin <brombin94@gmail.com>
- Pietro di Caprio <pietro@fabricators.ltd>
- תורמי Vanilla OS

## תרומה

תרומות מתקבלות בברכה! אנא ודא:

1. תלויויות מעודכנות עם `go get`, `go mod tidy`, ו-`go mod vendor`
2. הקוד עוקב אחר קונבנציות Go
3. תרגומים מעודכנים ב-Weblate

---

# Vanilla OS - Apx Package Manager

Apx is the default package manager for Vanilla OS. It is a wrapper around multiple package managers to install packages and run commands inside a managed container.

## Overview

**Apx** (/à·peks/) provides a unified interface for managing packages across different Linux distributions. It uses [distrobox](https://github.com/89luca89/distrobox) to create containerized environments where you can install packages without affecting your host system.

## Features

- **Multi-distribution support**: Install packages from different Linux distributions
- **Container-based**: Packages are installed in isolated containers
- **Package manager abstraction**: Unified interface for different package managers (apt, dnf, pacman, etc.)
- **Desktop integration**: Automatic desktop entry export for installed applications
- **Subsystem management**: Create and manage multiple subsystems with different stacks

## Stack

- **Language**: Go 1.25.0
- **Framework**: Vanilla OS SDK + CLI builder
- **UI**: Charmbracelet (bubbletea, lipgloss) for terminal UI
- **Container**: distrobox integration

## Project Structure

```
cmd/
  main.go              # Application entry point
  main_check.go        # String validation variant
  locales/             # Multi-language support (33+ languages)

internal/cli/
  structs.go           # CLI command definitions
  runtime.go           # Subsystem execution handlers
  subsystems.go        # Subsystem lifecycle management
  stacks.go            # Stack operations
  pkgmanagers.go       # Package manager operations

core/                  # Core business logic (subsystems, stacks, pkg management)
config/                # Configuration files
distrobox/             # distrobox integration scripts
```

## Building

### Prerequisites

- Go 1.25.0 or higher
- make
- podman or docker
- git

### Build Steps

```bash
# Clone and enter directory
git clone --recursive https://github.com/yourusername/Vanilla-.git
cd Vanilla-

# Build
make build

# Install system-wide
sudo make install
sudo make install-manpages

# Install to custom location
make install PREFIX=$HOME/.local
make install-manpages PREFIX=$HOME/.local
```

## Usage

### Basic Commands

```bash
# List available subsystems
apx subsystems list

# Create a new subsystem
apx subsystems new --name=ubuntu-latest

# Enter a subsystem
apx ubuntu-latest enter

# Install packages in subsystem
apx ubuntu-latest install package-name

# Run a command in subsystem
apx ubuntu-latest run command

# List available stacks
apx stacks list

# Manage package managers
apx pkgmanagers list
```

## Remote Access

Apx supports remote access via SSH, enabling package management and command execution on remote servers and devices.

### SSH Setup

```bash
# Connect to remote server
ssh user@remote-host

# Run Apx commands remotely
ssh user@remote-host "apx subsystems list"

# Install packages via SSH
ssh user@remote-host "apx ubuntu-latest install package-name"
```

### Using iPhone and Mobile Devices

For remote access from iPhone or other mobile devices, you can use SSH clients such as:

- **Termius** - Full-featured SSH client with container support
- **iSH** - SSH console for iOS
- **Prompt** - Professional SSH client for iOS

```bash
# Example: Connect from iPhone via Termius
ssh user@your-server.com
apx ubuntu-latest enter
```

### Secure Remote Access

```bash
# Setup public SSH key for secure access
ssh-copy-id -i ~/.ssh/id_rsa.pub user@remote-host

# Login without password
ssh user@remote-host "apx subsystems list"

# Execute remote script
ssh user@remote-host < /path/to/local/script.sh
```

## Android Support

Apx supports package management and command execution from Android devices via Termux and SSH.

### Termux Integration

Termux is a full-featured terminal for Android that supports running Apx directly.

```bash
# Install Termux from F-Droid
# https://f-droid.org/en/packages/com.termux/

# Install Apx in Termux
pkg install go git
git clone https://github.com/yourusername/Vanilla-.git
cd Vanilla-
make build
sudo make install
```

### Access Apx from Android

```bash
# Connect to your Apx server from Android
ssh user@your-server.com

# Run Apx commands remotely
ssh user@your-server.com "apx subsystems list"

# Install packages from Android
ssh user@your-server.com "apx ubuntu-latest install python3"
```

### Android SSH Applications

Popular SSH applications for Android:

- **Termux** - Direct command line and SSH
- **Termius** - Professional SSH client
- **JuiceSSH** - Enhanced SSH interface
- **ConnectBot** - Lightweight SSH client

```bash
# Example: Connect via Termius
1. Open Termius
2. Create new SSH connection
3. Enter your server details
4. Run: apx subsystems list
```

### Android App Development with Apx

```bash
# Install Android development tools in a subsystem
apx ubuntu-latest install android-sdk android-ndk

# Compile applications via SSH from Android
ssh user@dev-server.com "apx android-dev compile --project /path/to/project"

# Work with ADB remotely
ssh user@dev-server.com "apx android-dev adb shell"
```

### File Synchronization via Android

```bash
# Upload files from Android to Apx server
scp /sdcard/myproject/* user@server.com:/home/user/projects/

# Download results from Apx to Android
scp user@server.com:/build/output/* /sdcard/Downloads/
```

## AI and Artificial Intelligence

Apx integrates artificial intelligence capabilities to enhance package management and installation experience.

### Smart Package Recommendations

Apx uses machine learning to recommend relevant packages based on:
- Your previous installation history
- Common usage patterns
- Dependencies and similar projects

```bash
# Get AI recommendations for a package
apx ai-recommend package-name

# Analyze package with AI
apx ai-analyze ubuntu-latest
```

### AI-Powered Troubleshooting

AI can assist in resolving errors and creating an optimized environment:

```bash
# Get AI diagnosis for installation issues
apx ai-diagnose

# Get AI suggestions for troubleshooting
apx ai-troubleshoot error-message

# Get AI advice for optimal setup
apx ai-optimize
```

### Smart Dependency Resolution

AI analyzes package dependencies and finds:
- Missing dependencies
- Version conflicts
- Redundant packages

```bash
# Check dependencies with AI
apx ai-check-deps subsystem-name

# Get suggestions to remove redundant packages
apx ai-cleanup-deps subsystem-name

# Get AI report on system health
apx ai-health-report
```

### Multi-Language Support

AI supports Hebrew, English, and 30+ additional languages:

```bash
# Change language to Hebrew
apx --language=he ai-recommend package-name

# Get AI help in your language
apx ai-help --lang=en
```

## Dependencies

### Direct Dependencies
- `github.com/google/uuid` - UUID generation
- `github.com/vanilla-os/sdk` - Application framework
- `gopkg.in/yaml.v2` - YAML configuration

### UI Dependencies
- `github.com/charmbracelet/bubbletea` - Terminal UI framework
- `github.com/charmbracelet/lipgloss` - Terminal styling

### AI Dependencies
- `github.com/tmc/langchain-go` - LLM integration
- `github.com/sashabaranov/go-openai` - OpenAI integration (optional)
- `github.com/google/generative-ai-python` - Google AI integration (optional)

## Documentation

For detailed documentation, visit: https://docs.vanillaos.org/docs/en/apx

## Translations

Contribute translations via [Weblate](https://hosted.weblate.org/projects/vanilla-os/apx)

## License

GNU General Public License v3.0

## Authors

- Mirko Brombin <brombin94@gmail.com>
- Pietro di Caprio <pietro@fabricators.ltd>
- Vanilla OS Contributors

## Contributing

Contributions are welcome! Please ensure:

1. Dependencies are updated with `go get`, `go mod tidy`, and `go mod vendor`
2. Code follows Go conventions
3. Translations are updated in Weblate
