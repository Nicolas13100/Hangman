
# On supprime la tache avec le nom deja existante
Unregister-ScheduledTask -TaskName "OuvrirNavigateur"

# Définir le nom de la tâche planifiée
$taskName = "OuvrirNavigateur"

# Définir le chemin de l'exécutable de votre navigateur
$browserPath = "C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe"

# Créer une action pour ouvrir le navigateur
$action = New-ScheduledTaskAction -Execute $browserPath

# Créer un déclencheur pour exécuter la tâche toutes les 5 minutes
$trigger = New-ScheduledTaskTrigger -Once ` -At (Get-Date) -RepetitionInterval (New-TimeSpan -Minutes 5)

# Créer une nouvelle tâche planifiée
Register-ScheduledTask -TaskName $taskName -Action $action -Trigger $trigger

# On la désactive comme sa elle ne se lance pas auto
Disable-ScheduledTask "OuvrirNavigateur"

Write-Host "Tâche planifiée créée avec succès."

# Lancez la tâche
Enable-ScheduledTask "OuvrirNavigateur"


# Affichez les informations de la tache plannifiée
Get-ScheduledTaskInfo "OuvrirNavigateur"

# Désactivez la tache
Disable-ScheduledTask "OuvrirNavigateur"

# Modifiez la et lancer la tache totues les 10 mins
# crée un nouveau trigger
$trigger = New-ScheduledTaskTrigger -Once ` -At (Get-Date) -RepetitionInterval (New-TimeSpan -Minutes 10)

# On modifie le trigger
Set-ScheduledTask -TaskName "OuvrirNavigateur" -Trigger $trigger

# On relance la Task
Enable-ScheduledTask "OuvrirNavigateur"

